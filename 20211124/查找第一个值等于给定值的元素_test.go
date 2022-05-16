package _0211124

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  查找第一个值等于给定值的元素_test
 * @Date: 2021/11/24 14:20
 */

//@tag [二分查找]
func binarySearch_1(list []int,value int) int {
	low,high := 0,len(list) - 1
	for low <= high {
		mid := low + ((high - low) >> 1) // 数学上和 （high + low） / 2 相等
		if list[mid] > value {
			high = mid - 1
		}else if list[mid] < value {
			low = mid + 1
		}else if mid == 0 || list[mid - 1] != value {
			return mid
		}else {
			high = mid - 1
		}
	}
	return -1
}


func Test_1(t *testing.T) {
	nums := []int {1,3,4,5,6,8,8,8,11,18}
	fmt.Println(binarySearch_1(nums,8))

}