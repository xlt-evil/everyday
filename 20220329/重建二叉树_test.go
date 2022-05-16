package _0220329

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  重建二叉树_test
 * @Date: 2022/03/29 9:28
 */
/**
 * Definition for a binary tree node.
 */

 type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

//https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	i := 0 // 根节点的索引
	root := &TreeNode{Val: preorder[0]}
	for  ;i < len(inorder);i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1],inorder[:i]) // 对于左子树来说前序遍历的索引为1的元素总是左子树的开始,i为中序遍历的根节点索引，表示左子树的元素数量
	root.Right = buildTree(preorder[i+1:],inorder[i+1:]) // 前序和中序 他们最终根+左的元素是一样多的，所以有明确的右子树边界
	return root
}

 func Test_buildTree(t *testing.T) {
 	buildTree([]int{3,9,20,15,7},[]int{9,3,15,20,7})
 }