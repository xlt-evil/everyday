package _0211210

/**
 * @Author: hxy
 * @Description:
 * @File:  在排序数组中查找数字 I_test
 * @Date: 2021/12/10 11:53
 */

func search(nums []int, target int) int {
	for i := 0; i < len(nums);i++ {
		if nums[i] == target {
			j := i + 1
			for ;j < len(nums) &&nums[j] == target;{
				j++
			}
			return j - i
		}
	}
	return -1
}