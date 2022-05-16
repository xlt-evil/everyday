package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉搜索树中的搜索_test
 * @Date: 2021/11/18 15:02
 */
//@tag [二叉树]
func searchBST(root *TreeNode, val int) *TreeNode {
	var (
		dfs func(*TreeNode) *TreeNode
	)
	
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == val {
			return node
		}
		if n := dfs(node.Left);n != nil {
			return n
		}
		if n := dfs(node.Right);n != nil {
			return n
		}
		return nil
	}
	return dfs(root)
}