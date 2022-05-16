package _0220209

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  syncPool
 * @Date: 2022/02/09 10:08
 */

type Student struct {
	Name string
}

func (s *Student) reset() {
	s.Name = ""
}

// 连接池获取对象的顺序从当前的private / share 获取，再从其他p的share / victim 获取都没有new 一个
func SyncPool() {
	p := sync.Pool{
		New: func() interface{} {
			return new(Student)
		},
	}
	for i := 0;i < 100;i++ {
		go func(i int) {
			s := p.Get().(*Student)
			s.reset()
			s.Name = s.Name + "," + strconv.Itoa(i) // 那出来的对象需要reset一下，避免垃圾数据污染
			fmt.Println(s.Name)

		}(i)
	}
	c := time.NewTimer(10 * time.Second)
	select {
	case <-c.C:
		
	}
}

func Test_syncPool(t *testing.T) {
	SyncPool()
}