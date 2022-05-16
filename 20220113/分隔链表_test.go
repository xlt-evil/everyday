package _0220113

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  分隔链表_test
 * @Date: 2022/01/13 16:13
 */


//@tag [链表]
func partition(head *ListNode, x int) *ListNode {
	var (
		firstBig, _,firstSmall,cur *ListNode
	)
	cur = head
	_ = head
	for cur != nil {
		if cur.Val >= x && firstBig == nil{
			firstBig = cur
		}else {
			_ = cur
			if firstSmall == nil {
				firstSmall = cur
			}
		}
		cur = cur.Next
	}
	return firstSmall
}


func Test_partition(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 2}
	n1.Next = n2
	n2.Next = n3
	ListNodeTraverse(partition(n1,3))
}