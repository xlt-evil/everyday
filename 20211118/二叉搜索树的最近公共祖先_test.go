package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉搜索树的最近公共祖先_test
 * @Date: 2021/11/18 16:57
 */
//@tag [二叉树]
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root
	for {
		if p.Val > node.Val && q.Val > node.Val {
			node = node.Right
		}else if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		}else {
			break
		}
	}
	return node
}
