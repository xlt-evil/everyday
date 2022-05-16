package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉树的后序遍历_test
 * @Date: 2021/11/18 13:35
 */
//@tag [二叉树]
// 左右根
func postorderTraversal(root *TreeNode) []int {
	var (
		result []int
		dfs func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result,node.Val)
	}
	dfs(root)
	return result
}