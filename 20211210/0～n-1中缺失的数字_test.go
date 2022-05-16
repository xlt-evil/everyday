package _0211210

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  0～n-1中缺失的数字_test
 * @Date: 2021/12/10 13:40
 */

//@tag [双指针]
func missingNumber(nums []int) int {
	low,high := 0,len(nums) - 1
	for low <= high {
		mid := low + ((high - low)>>1)
		if nums[mid] == mid {
			low  = mid + 1
		}else {
			high = mid - 1
		}
	}
	return low
}

func Test_missingNumber(t *testing.T) {
	fmt.Println(missingNumber([]int{0,1,2}))
}


