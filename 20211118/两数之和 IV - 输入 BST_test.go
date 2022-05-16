package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  两数之和 IV - 输入 BST_test
 * @Date: 2021/11/18 16:40
 */

//@tag [二叉树]
func findTarget(root *TreeNode, k int) bool {
	var (
		m = make(map[int]struct{})
		dfs func(node *TreeNode) bool
	)
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _,ok := m[k - node.Val];ok {
			return true
		}else {
			m[node.Val] = struct{}{}
		}
		return dfs(node.Right) || dfs(node.Left)
	}
	return dfs(root)
}