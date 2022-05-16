package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉搜索树中的插入操作_test
 * @Date: 2021/11/18 15:07
 */
//@tag [二叉树]
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode)  {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				return
			}
			dfs(node.Left)
			return
		}else {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				return
			}
			dfs(node.Right)
			return
		}
	}
	dfs(root)
	return root
}