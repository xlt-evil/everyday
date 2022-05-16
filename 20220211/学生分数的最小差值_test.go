package _0220211

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  学生分数的最小差值_test
 * @Date: 2022/02/11 14:33
 */

//@tag 滑动窗口
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := nums[k - 1] - nums[0]
	for i := k;i < len(nums);i++ {
		ans = min(ans, nums[i] - nums[i - k + 1])
	}
	return ans

}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func Test_minimumDifference(t *testing.T) {
	fmt.Println(minimumDifference([]int{9,4,1,7},2))
}