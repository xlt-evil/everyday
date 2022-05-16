package _0211102

/**
 * @Author: hxy
 * @Description:
 * @File:  删除节点_test
 * @Date: 2021/11/02 9:46
 */

//@tag [链表,递归]
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	pre := node
	for node.Next != nil {
		pre = node
		node.Val = node.Next.Val
		node = node.Next
	}
	pre.Next = nil
	return
}