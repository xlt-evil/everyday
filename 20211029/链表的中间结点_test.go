package _0211029

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  链表的中间结点_test
 * @Date: 2021/11/01 14:10
 */

//@tag [链表,快慢指针]
func MiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow,fast := head,head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}


func Test_MiddleNode(t *testing.T) {
	l1 := CreateListNode([]int{
		1,2,3,4,5,
	})
	l1 = MiddleNode(l1)
	ListNodeTraverse(l1)
}