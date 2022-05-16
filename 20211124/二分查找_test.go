package _0211124

/**
 * @Author: hxy
 * @Description:
 * @File:  二分查找_test
 * @Date: 2021/11/24 15:40
 */
//@tag [二分查找]
func search(nums []int, target int) int {
	low,high := 0 ,len(nums) - 1
	for low <= high {
		mid := low + ((high - low)>>1)
		if nums[mid] > target {
			high = mid - 1
		}else if nums[mid] < target {
			low = mid + 1
		}else {
			return mid
		}
	}
	return -1
}