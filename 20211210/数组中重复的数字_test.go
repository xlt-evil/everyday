package _0211210

import "sort"

/**
 * @Author: hxy
 * @Description:
 * @File:  数组中重复的数字_test
 * @Date: 2021/12/10 11:50
 */


func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0 ;i < len(nums) - 1;i++ {
		if nums[i] == nums[i + 1] {
			return nums[i]
		}
	}
	return -1
}