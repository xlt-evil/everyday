package _0211209

/**
 * @Author: hxy
 * @Description:
 * @File:  从尾到头打印链表_test
 * @Date: 2021/12/09 16:50
 */

func reversePrint(head *ListNode) []int {
	var (
		ans []int
		dfs func(*ListNode)
	)
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next)
		ans = append(ans,node.Val)
	}
	dfs(head)
	return ans
}