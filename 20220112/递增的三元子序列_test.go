package _0220112

import "math"

/**
 * @Author: hxy
 * @Description:
 * @File:  递增的三元子序列_test
 * @Date: 2022/01/12 10:06
 */


// 贪心算法
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first,second := nums[0],math.MaxInt32
	for i := 1; i < len(nums);i++{
		if nums[i] > second {
			return true
		}else if nums[i] > first {
			second = nums[i]
		}else 	{
			first = nums[i]
		}
	}
	return false
}

