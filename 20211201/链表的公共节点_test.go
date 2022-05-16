package _0211201

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  链表的公共节点_test
 * @Date: 2021/12/01 13:39
 */

func ListNodeFather(n1,n2 *ListNode) *ListNode {
	p1,p2 := n1,n2
	l1,l2 := GetLength(n1),GetLength(n2)
	for i := 0;i < l1 - l2;i++ {
		p1 = p1.Next
	}
	for i := 0 ;i < l2 - l1;i++ {
		p2 = p2.Next
	}
	for p1 != nil && p2 != nil {
		if p1.Val == p2.Val {
			return p1
		}
		p1,p2 = p1.Next,p2.Next
	}
	return nil
}

func GetLength(n1 *ListNode) int {
	length := 0
	for n1 != nil {
		length++
		n1 = n1.Next
	}
	return length
}

type ss struct {
	Val int `json:"val"`
}

func Test_ListNodeFather(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n6

	n4.Next = n5
	n5.Next = n6

	n6.Next = n7

	fmt.Println(ListNodeFather(n1,n4).Val)
}
