package _0211201

/**
 * @Author: hxy
 * @Description:
 * @File:  填充每个节点的下一个右侧节点指针_test
 * @Date: 2021/12/01 10:16
 */
//@tag [广度优先搜索,二叉树]
func connect(root *Node) *Node {
	q := []*Node{root}
	for len(q) != 0 {
		length := len(q)
		for i := 0 ;i < length ;i++ {
			if i < length - 1 {
				q[i].Next = q[i + 1].Next
			}
			if q[i].Left != nil {
				q = append(q,q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q,q[i].Right)
			}
		}
		q = q[length:]
	}
	return root
}