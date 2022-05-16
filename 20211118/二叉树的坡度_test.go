package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉树的坡度_test
 * @Date: 2021/11/18 9:48
 */

/**
 * Definition for a binary tree node.
 */

//@tag [二叉树]
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func findTilt(root *TreeNode) int {
	var (
		ant int
		dfs func(*TreeNode) int
	)

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		ant += abs(l - r)
		return l + r + node.Val
	}
	dfs(root)
	return ant
}

