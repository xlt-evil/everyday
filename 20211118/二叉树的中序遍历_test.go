package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉树的中序遍历_test
 * @Date: 2021/11/18 11:22
 */
// 左根右
//@tag [二叉树]
func inorderTraversal(root *TreeNode) []int {
	var (
		result []int
		dfs func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result,node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return result
}