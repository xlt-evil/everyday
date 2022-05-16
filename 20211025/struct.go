package _20211025

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

