package _0211124

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  查找最后一个小于等于给定值的元素_test
 * @Date: 2021/11/24 14:58
 */
//@tag [二分查找]
func binarySearch_4(list []int,value int) int {
	low,high,mid := 0 ,len(list) - 1,-1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if list[mid] <= value {
			if mid == len(list) || list[mid+1] > value {
				return mid
			}else {
				low = mid + 1
			}
		}else {
			high = mid - 1
		}
	}
	return -1
}

func Test_4(t *testing.T) {
	nums := []int {1,3,4,5,6,8,8,8,11,18}
	fmt.Println(nums[binarySearch_4(nums,7)])
}