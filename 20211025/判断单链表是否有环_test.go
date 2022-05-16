package _20211025

/**
 * @Author: hxy
 * @Description:
 * @File:  判断单链表是否有环_test
 * @Date: 2021/10/25 15:26
 */

//@tag [链表,快慢指针]
func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow,quick := head,head
	for quick != nil {
		if quick.Next == nil {
			return false
		}
		quick = quick.Next.Next
		if quick == slow {
			return true
		}
		slow = slow.Next
	}
	return false
}