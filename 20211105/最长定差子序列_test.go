package _0211105

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  最长定差子序列_test
 * @Date: 2021/11/05 17:22
 */
//@tag [动态规划]
func longestSubsequence(arr []int, difference int) (ans int) {
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return
}

func Test_longestSubsequence(t *testing.T) {
	//arr = [1,2,3,4], difference = 1
	fmt.Println(longestSubsequence([]int{1,3,5,7},2))
}


