package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  对称二叉树_test
 * @Date: 2021/11/18 14:03
 */
//@tag [二叉树]
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Right,root.Left)
}

func check(r,l *TreeNode) bool {
	if r == nil && l != nil {
		return false
	}
	if r != nil && l == nil {
		return false
	}
	if r == nil && l == nil {
		return true
	}
	return r.Val == l.Val && check(r.Right,l.Left) && check(r.Left,l.Right)
}