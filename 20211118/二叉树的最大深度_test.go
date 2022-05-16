package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉树的最大深度_test
 * @Date: 2021/11/18 13:51
 */
//@tag [二叉树]
func maxDepth(root *TreeNode) int {
	var (
		dfs  func(*TreeNode) int
	)
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(dfs(node.Left),dfs(node.Right)) + 1
	}
	return dfs(root)
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}