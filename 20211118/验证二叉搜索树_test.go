package _0211118

import (
	"fmt"
	"math"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  验证二叉搜索树_test
 * @Date: 2021/11/18 15:25
 */
//@tag [二叉树]
func isValidBST(root *TreeNode) bool {
	var (
		dfs func (*TreeNode,int, int) bool
	)
	dfs = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		return dfs(node.Left,lower,node.Val) && dfs(node.Right,node.Val,upper)
	}
	return dfs(root,math.MinInt64,math.MaxInt64)
}

func Test_isValidBST(t *testing.T)  {
	ts := &TreeNode{Val: 5,Left: &TreeNode{Val: 3,Left: &TreeNode{Val: 2,Right: &TreeNode{Val: 10}}}}
	fmt.Println(isValidBST(ts))
}

