package _0211102

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  最大子序和_test
 * @Date: 2021/11/02 10:52
 */

//@tag [动态规划]
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	dp,m := nums[0],nums[0]
	// f(i) = max(f(i - 1) + num(i), num(i))
	for i := 1;i < len(nums);i++ {
		//前面的积蓄是正的资产是有意义的，就加起来，最后结果会更大，不然就是重新开始。
		//当前求的是当子序列到i时的最大合是多少
		dp = max(dp + nums[i], nums[i])
		// 全局的结果
		m = max(dp,m)
	}
	return m
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_maxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}