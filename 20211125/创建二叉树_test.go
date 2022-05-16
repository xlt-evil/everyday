package _0211125

import (
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  创建二叉树
 * @Date: 2021/11/25 14:38
 */

//@tag [二叉树]
//用前序的方式创建二叉树
func CreateBinaryTree(result []int) *TreeNode {
	var (
		dfs func() *TreeNode
	)
	dfs = func() *TreeNode {
		if len(result) == 0 {
			return nil
		}
		val := result[0]
		result = result[1:]
		var cur *TreeNode
		if val != 0 {
			cur = new(TreeNode)
			cur.Val = val
			cur.Left = dfs()
			cur.Right = dfs()
		}
		return cur
	}
	return dfs()
}


func Test_CreateBinaryTree(t *testing.T) {
	//nums := []int{10,9,8,7,6,0,11,0,12}

}
