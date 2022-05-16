package _20211025

import (
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  创建链表_test
 * @Date: 2021/10/25 14:53
 */


//@tag [链表]
func CreateListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	head := new(ListNode)
	cur := head

	for _,v := range list {
		cur.Next = &ListNode{Val: v,Next: nil}
		cur = cur.Next
	}

	return head.Next
}

func Test_CreateListNode(t *testing.T) {
	list := []int{1,2,3,4,5,6}
	l := CreateListNode(list)
	ListNodeTraverse(l)
}