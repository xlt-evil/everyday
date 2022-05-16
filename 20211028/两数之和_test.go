package _0211028

/**
 * @Author: hxy
 * @Description:
 * @File:  两数之和_test
 * @Date: 2021/12/03 9:50
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry int
		sum int
		head = new(ListNode)
	)
	p := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = new(ListNode)
		}
		if l2 == nil {
			l2 = new(ListNode)
		}
		sum = l1.Val + l2.Val + carry
		carry = 0
		if sum / 10 != 0 {
			carry = 1
			sum = sum % 10
		}
		t := new(ListNode)
		t.Val = sum
		p.Next = t
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry != 0 {
		p.Next = &ListNode{Val:1}
	}
	return head.Next
}