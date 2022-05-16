package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉树的前序遍历_test
 * @Date: 2021/11/18 11:13
 */
//@tag [二叉树]
// 根左右
func preorderTraversal(root *TreeNode) []int {
	var (
		result []int
		dfs func(node *TreeNode)
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result,node.Val)
		dfs(node.Left)
		dfs(node.Right)
		return
	}
	dfs(root)
	return result
}