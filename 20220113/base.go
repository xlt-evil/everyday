package _0220113

import "fmt"

/**
 * @Author: hxy
 * @Description:
 * @File:  base
 * @Date: 2022/01/13 16:14
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