package _0220329

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  青蛙跳台阶问题_test
 * @Date: 2022/03/29 10:16
 */

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n < 3 {
		return n
	}
	mod := 1000000007
	dp1,dp2 := 1,2
	sum := 0
	for i := 3;i <= n;i++ {
		sum = (dp1 + dp2) % mod
		dp1,dp2 = dp2,sum
	}
	return sum
}

func Test_numWays(t *testing.T) {
	fmt.Println(numWays(7))
}
