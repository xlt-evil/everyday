package _0211102

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  存在重复元素_test
 * @Date: 2021/11/02 10:46
 */
//@tag [数组]
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0 ;i < len(nums);i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			return true
		}
	}
	return false
}

func Test_containsDuplicate(t *testing.T) {
	fmt.Println(containsDuplicate([]int{1,3,4,5,1}))
}