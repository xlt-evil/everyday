package _0211209

/**
 * @Author: hxy
 * @Description:
 * @File:  复杂链表的复制_test
 * @Date: 2021/12/09 17:28
 */

// @tag [递归]
func copyRandomList(head *Node) *Node {
	var (
		record = map[*Node]*Node{}
		dfs func(node *Node) *Node
	)
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n,ok := record[node];ok {
			return n
		}
		newNode := &Node{Val: node.Val}
		record[node] = newNode
		newNode.Next = dfs(node.Next)
		newNode.Random = dfs(node.Random)
		return newNode
	}
	return dfs(head)
}