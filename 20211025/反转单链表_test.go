package _20211025

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  反转单链表
 * @Date: 2021/10/25 14:15
 */



//@tag [链表,双指针]
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var cur,pre,next *ListNode

	cur = head

	for cur != nil {
		next = cur
		cur = cur.Next
		next.Next = pre
		pre = next
	}
	return next
}

func reverse(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	head := reverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return head
}


func Test_ReverseList(t *testing.T) {
	list := CreateListNode([]int{1,2,3,4,5,6})
	list = reverse(list)
	ListNodeTraverse(list)
}