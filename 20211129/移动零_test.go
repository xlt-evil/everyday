package _0211129

/**
 * @Author: hxy
 * @Description:
 * @File:  移动零_test
 * @Date: 2021/11/29 13:43
 */

func moveZeroes(nums []int)  {
	if len(nums) == 0 {
		return
	}
	right,left := 0,0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left],nums[right] = nums[right],nums[left]
			left++
		}
		right++
	}
}