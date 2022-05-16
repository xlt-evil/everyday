package _0211118

/**
 * @Author: hxy
 * @Description:
 * @File:  二叉树的层序遍历_test
 * @Date: 2021/11/18 13:39
 */
//@tag [二叉树]
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	q := make([]*TreeNode,1)
	q[0] = root
	for len(q) != 0 {
		length := len(q)
		cur := make([]int,0)
		 for i := 0;i < length;i++ {
		 	if q[i].Left != nil {
		 		q = append(q,q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q,q[i].Right)
			}
			cur = append(cur,q[i].Val)
		 }
	    result = append(result,cur)
		q = q[length:]
	}
	return result
}