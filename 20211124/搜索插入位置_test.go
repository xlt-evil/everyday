package _0211124

/**
 * @Author: hxy
 * @Description:
 * @File:  搜索插入位置_test
 * @Date: 2021/11/24 15:51
 */

//@tag [二分查找]

// 思路是通过查找最后一个小于当前的数的小标 +1 如果都不小于就是 0
func searchInsert(nums []int, target int) int {
	low,high := 0 ,len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] < target {
			if mid == len(nums) - 1 || nums[mid + 1] > target {
				return mid + 1
			}else {
				low = mid + 1
			}
		}else {
			high = mid - 1
		}
	}
	return 0
}