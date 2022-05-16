package _0211029

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  删除倒数第N个节点
 * @Date: 2021/11/01 14:07
 */

//@tag [链表,快慢指针]
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{Next:head}
	slow,fast := p,head
	for i := 0; i <  n;i++ {
		fast = fast.Next
	}

	for ;fast != nil; {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return p.Next
}

func Test_RemoveNthFromEnd(t *testing.T) {
	l := CreateListNode([]int{1,2,3,4,5})
	l = RemoveNthFromEnd(l,2)
	ListNodeTraverse(l)
}