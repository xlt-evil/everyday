package _0211201

/**
 * @Author: hxy
 * @Description:
 * @File:  base
 * @Date: 2021/11/30 11:42
 */

func check(m map[[2]int]bool,i,j int) bool {
	return m[[2]int{i,j}]
}

func mark(m map[[2]int]bool,i,j int) {
	m[[2]int{i,j}] = true
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

 type Node struct {
     Val int
     Left *Node
     Right *Node
     Next *Node
}

type ListNode struct {
	Val int
	Next *ListNode
}