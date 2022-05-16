package _0211115

/**
 * @Author: hxy
 * @Description:
 * @File:  两个数组的交集 II_test
 * @Date: 2021/11/15 14:16
 */

//@tag [哈希]
func intersect(nums1 []int, nums2 []int) []int {
	answer := []int{}
	m := make(map[int]int,0)
	for _,v := range nums1 {
		m[v]++
	}
	for _,v := range nums2 {
		if x,_ := m[v];x > 0 {
			answer = append(answer,v)
			m[v]--
		}
	}
	return answer
}