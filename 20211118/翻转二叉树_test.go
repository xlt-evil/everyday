package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  翻转二叉树_test
 * @Date: 2021/11/18 14:15
 */
//@tag [二叉树]
func invertTree(root *TreeNode) *TreeNode {
	var (
		dfs func(node *TreeNode) *TreeNode
	)
	dfs = func(node *TreeNode) *TreeNode{
		if node == nil {
			return nil
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		node.Left = right
		node.Right = left
		return node
	}
	return dfs(root)
}