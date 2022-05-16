package _0211203

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  构建二叉树
 * @Date: 2021/12/03 17:38
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 跟左右
// 左跟右
//@tag [二叉树]
func buildTree(preorder []int,inorder []int) *TreeNode {
	if len(preorder) == 0{
		return nil
	}
	// 1. 创建根节点
	root := &TreeNode{Val: preorder[0]}
	// 2. 获取根节点在中序遍历数组中的index
	var i int
	for index,value := range inorder{
		if value == preorder[0]{
			i = index
			break
		}
	}
	// 3. 递归
	root.Left = buildTree(preorder[1:i+1],inorder[:i])
	root.Right = buildTree(preorder[i+1:],inorder[i+1:])
	return root
}

func Test_buildTree(t *testing.T) {
	n := buildTree([]int{3,9,20,15,7},[]int{9,3,15,20,7})
	fmt.Println(n)
}