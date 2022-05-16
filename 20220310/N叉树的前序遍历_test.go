package _0220310

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  N叉树的前序遍历
 * @Date: 2022/03/10 10:55
 */

type Node struct {
     Val int
     Children []*Node
}

//https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
func preorder(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, ch := range node.Children {
			dfs(ch)
		}
	}
	dfs(root)
	return
}


func Test_preorder(t *testing.T) {

}