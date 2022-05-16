package _0211203

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  两个链表生成相加链表_test
 * @Date: 2021/12/03 9:37
 */
func addInList( head1 *ListNode ,  head2 *ListNode ) *ListNode {
	p1 := ReverseList(head1)
	p2 := ReverseList(head2)
	p3 := addTwoNumbers(p1,p2)
	return ReverseList(p3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry int
		sum int
		head = new(ListNode)
	)
	p := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = new(ListNode)
		}
		if l2 == nil {
			l2 = new(ListNode)
		}
		sum = l1.Val + l2.Val + carry
		carry = 0
		if sum / 10 != 0 {
			carry = 1
			sum = sum % 10
		}
		t := new(ListNode)
		t.Val = sum
		p.Next = t
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry != 0 {
		p.Next = &ListNode{Val:1}
	}
	return head.Next
}

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


func Test_addInList(t *testing.T) {
	n1 := &ListNode{Val: 9}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3


	n4 := &ListNode{Val: 6}
	n5 := &ListNode{Val: 3}
	n4.Next = n5

	p := addInList(n1,n4)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}