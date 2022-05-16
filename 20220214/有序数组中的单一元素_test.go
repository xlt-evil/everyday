package _0220214

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  有序数组中的单一元素_test
 * @Date: 2022/02/14 14:05
 */

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		fmt.Println(mid)
		// mid ^ 1
		// mid 是奇数 = mid - 1
		// mid 是偶数 = mid + 1
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func Test_single(t *testing.T) {
	fmt.Println(singleNonDuplicate([]int{1,1,3,3,2,4,4,5,5}))
}