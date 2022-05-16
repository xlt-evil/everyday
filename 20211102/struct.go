package _0211102

import "fmt"

/**
 * @Author: hxy
 * @Description:
 * @File:  struct
 * @Date: 2021/10/25 14:26
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func ListNodeTraverse(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

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