// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mvcc

import (
	"sort"
	"sync"

	"github.com/google/btree"
	"go.uber.org/zap"
)

type index interface {
	Get(key []byte, atRev int64) (rev, created revision, ver int64, err error)
	Range(key, end []byte, atRev int64) ([][]byte, []revision)
	Revisions(key, end []byte, atRev int64) []revision
	Put(key []byte, rev revision)
	Tombstone(key []byte, rev revision) error
	RangeSince(key, end []byte, rev int64) []revision
	Compact(rev int64) map[revision]struct{}
	Keep(rev int64) map[revision]struct{}
	Equal(b index) bool

	Insert(ki *keyIndex)
	KeyIndex(ki *keyIndex) *keyIndex
}

// 用 google 开源的 btree 的实现
// btree 的节点结构
type treeIndex struct {
	sync.RWMutex
	tree *btree.BTree
	lg   *zap.Logger
}

func newTreeIndex(lg *zap.Logger) index {
	return &treeIndex{
		tree: btree.New(32),
		lg:   lg,
	}
}

func (ti *treeIndex) Put(key []byte, rev revision) {
	// btree 的 key 是 keyIndex 结构的;
	keyi := &keyIndex{key: key}

	ti.Lock()
	defer ti.Unlock()
	// 通过 key 去 btree 里找到 keyIndex 结构体
	item := ti.tree.Get(keyi)
	if item == nil {
		// 如果 btree 树上不存在，那么就 insert 上去
		// 先添加一个 generation
		keyi.put(ti.lg, rev.main, rev.sub)
		// 添加 keyIndex 节点到 btree
		ti.tree.ReplaceOrInsert(keyi)
		return
	}
	// 找到了现成的，那么就更新这个 keyIndex
	okeyi := item.(*keyIndex)
	okeyi.put(ti.lg, rev.main, rev.sub)
}

// key: 用户的 key
// atRev : 本次获取的版本
func (ti *treeIndex) Get(key []byte, atRev int64) (modified, created revision, ver int64, err error) {
	keyi := &keyIndex{key: key}
	ti.RLock()
	defer ti.RUnlock()
	// 根据 key 获取到 keyIndex 结构（也就找到了 revision）
	if keyi = ti.keyIndex(keyi); keyi == nil {
		return revision{}, revision{}, 0, ErrRevisionNotFound
	}
	// 获取到 keyIndex 之后，就可以根据 revision 去查找对应的版本了
	return keyi.get(ti.lg, atRev)
}

func (ti *treeIndex) KeyIndex(keyi *keyIndex) *keyIndex {
	ti.RLock()
	defer ti.RUnlock()
	// 根据 key 查找 keyIndex
	return ti.keyIndex(keyi)
}

// 通过 key 去 b 树上查找 keyIndex 。keyIndex 的 Less 定义成比较 key
func (ti *treeIndex) keyIndex(keyi *keyIndex) *keyIndex {
	if item := ti.tree.Get(keyi); item != nil {
		return item.(*keyIndex)
	}
	return nil
}

// 遍历 b 树节点
func (ti *treeIndex) visit(key, end []byte, f func(ki *keyIndex)) {
	keyi, endi := &keyIndex{key: key}, &keyIndex{key: end}

	ti.RLock()
	defer ti.RUnlock()

	ti.tree.AscendGreaterOrEqual(keyi, func(item btree.Item) bool {
		if len(endi.key) > 0 && !item.Less(endi) {
			return false
		}
		f(item.(*keyIndex))
		return true
	})
}

func (ti *treeIndex) Revisions(key, end []byte, atRev int64) (revs []revision) {
	if end == nil {
		rev, _, _, err := ti.Get(key, atRev)
		if err != nil {
			return nil
		}
		return []revision{rev}
	}
	// 遍历 [key, end] 的 keyIndex
	ti.visit(key, end, func(ki *keyIndex) {
		// 获取到指定的 revision 结构
		if rev, _, _, err := ki.get(ti.lg, atRev); err == nil {
			revs = append(revs, rev)
		}
	})
	return revs
}

func (ti *treeIndex) Range(key, end []byte, atRev int64) (keys [][]byte, revs []revision) {
	if end == nil {
		rev, _, _, err := ti.Get(key, atRev)
		if err != nil {
			return nil, nil
		}
		return [][]byte{key}, []revision{rev}
	}
	// 遍历 [key, end]，获取到指定的 revision 和 key，返回对应的数组；
	ti.visit(key, end, func(ki *keyIndex) {
		if rev, _, _, err := ki.get(ti.lg, atRev); err == nil {
			revs = append(revs, rev)
			keys = append(keys, ki.key)
		}
	})
	return keys, revs
}

func (ti *treeIndex) Tombstone(key []byte, rev revision) error {
	keyi := &keyIndex{key: key}

	ti.Lock()
	defer ti.Unlock()
	item := ti.tree.Get(keyi)
	if item == nil {
		return ErrRevisionNotFound
	}

	ki := item.(*keyIndex)
	// 删除版本（添加一个特殊的东西）
	return ki.tombstone(ti.lg, rev.main, rev.sub)
}

// RangeSince returns all revisions from key(including) to end(excluding)
// at or after the given rev. The returned slice is sorted in the order
// of revision.
func (ti *treeIndex) RangeSince(key, end []byte, rev int64) []revision {
	keyi := &keyIndex{key: key}

	ti.RLock()
	defer ti.RUnlock()

	if end == nil {
		item := ti.tree.Get(keyi)
		if item == nil {
			return nil
		}
		keyi = item.(*keyIndex)
		return keyi.since(ti.lg, rev)
	}

	endi := &keyIndex{key: end}
	var revs []revision
	ti.tree.AscendGreaterOrEqual(keyi, func(item btree.Item) bool {
		if len(endi.key) > 0 && !item.Less(endi) {
			return false
		}
		curKeyi := item.(*keyIndex)
		revs = append(revs, curKeyi.since(ti.lg, rev)...)
		return true
	})
	sort.Sort(revisions(revs))

	return revs
}

// 压缩树占用空间，把一些无效的内容清理掉；
func (ti *treeIndex) Compact(rev int64) map[revision]struct{} {
	available := make(map[revision]struct{})
	if ti.lg != nil {
		ti.lg.Info("compact tree index", zap.Int64("revision", rev))
	} else {
		plog.Printf("store.index: compact %d", rev)
	}
	ti.Lock()
	clone := ti.tree.Clone()
	ti.Unlock()

	// 克隆出一个 b 树，用来安全的遍历
	clone.Ascend(func(item btree.Item) bool {
		keyi := item.(*keyIndex)
		//Lock is needed here to prevent modification to the keyIndex while
		//compaction is going on or revision added to empty before deletion
		ti.Lock()
		keyi.compact(ti.lg, rev, available)
		// 如果空了就从 b 树删除
		if keyi.isEmpty() {
			item := ti.tree.Delete(keyi)
			if item == nil {
				if ti.lg != nil {
					ti.lg.Panic("failed to delete during compaction")
				} else {
					plog.Panic("store.index: unexpected delete failure during compaction")
				}
			}
		}
		ti.Unlock()
		return true
	})
	return available
}

// Keep finds all revisions to be kept for a Compaction at the given rev.
func (ti *treeIndex) Keep(rev int64) map[revision]struct{} {
	available := make(map[revision]struct{})
	ti.RLock()
	defer ti.RUnlock()
	ti.tree.Ascend(func(i btree.Item) bool {
		keyi := i.(*keyIndex)
		keyi.keep(rev, available)
		return true
	})
	return available
}

func (ti *treeIndex) Equal(bi index) bool {
	b := bi.(*treeIndex)

	if ti.tree.Len() != b.tree.Len() {
		return false
	}

	equal := true

	ti.tree.Ascend(func(item btree.Item) bool {
		aki := item.(*keyIndex)
		bki := b.tree.Get(item).(*keyIndex)
		if !aki.equal(bki) {
			equal = false
			return false
		}
		return true
	})

	return equal
}

func (ti *treeIndex) Insert(ki *keyIndex) {
	ti.Lock()
	defer ti.Unlock()
	ti.tree.ReplaceOrInsert(ki)
}
