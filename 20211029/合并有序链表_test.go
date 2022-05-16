package _0211029

import (
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  合并有序链表_test
 * @Date: 2021/10/29 14:09
 */

//@tag [链表,双指针]
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := new(ListNode)
	p := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			sum := 0
			if l1.Val < l2.Val {
				sum = l1.Val
				l1 = l1.Next
			}else {
				sum = l2.Val
				l2 = l2.Next
			}
			p.Next = &ListNode{Val: sum}
			p = p.Next
			continue
		}
		if l1 == nil {
			p.Next = l2
			break
		}
		if l2 == nil {
			p.Next = l1
			break
		}
	}

	return head.Next
}

func Test_MergeTwoLists(t *testing.T) {
	l1 := CreateListNode([]int{1,2,4})
	l2 := CreateListNode([]int{1,3,4})
	l3 := MergeTwoLists(l1,l2)
	ListNodeTraverse(l3)
}