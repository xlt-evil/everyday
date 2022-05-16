package _0211208

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  斐波那契数列_test
 * @Date: 2021/12/08 10:22
 */

//@tag [递归]
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)
}

//@tag [动态规划]
func fibDP(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	dp1, dp2, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		dp1 = dp2
		dp2 = r
		r = (dp1 + dp2) % mod
	}
	return r
}

func Test_fib(t *testing.T) {
	// 134903163
	// 701408733
	fmt.Println(fibDP(5))
}

