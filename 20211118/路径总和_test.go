package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  路径总和_test
 * @Date: 2021/11/18 14:35
 */
//@tag [二叉树]
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum - root.Val == 0
	}
	return hasPathSum(root.Left,targetSum - root.Val) || hasPathSum(root.Right,targetSum - root.Val)
}