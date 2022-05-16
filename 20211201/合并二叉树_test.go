package _0211201

/**
 * @Author: hxy
 * @Description:
 * @File:  合并二叉树
 * @Date: 2021/12/01 10:12
 */
//@tag [递归,二叉树]
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left,root2.Left)
	root1.Right = mergeTrees(root1.Right,root2.Right)
	return root1
}