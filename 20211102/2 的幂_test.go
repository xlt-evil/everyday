package _0211102

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  2 的幂_test
 * @Date: 2021/11/02 15:18
 */

//@tag [递归]

// 递归
// 1. 问题可以拆成子问题
// 2. 子问题，除了数据规模，求解方式一致
// 3. 存在边界
func isPowerOfTwo(n int) bool {
	return check(1,n,2)
}

func check (start int,n int,power int) bool {
	if start == n {
		return true
	}
	if start > n {
		return false
	}
	return check(start * power,n,power)
}

func Test_isPowerOfTwo(t *testing.T) {
	fmt.Println(isPowerOfTwo(10))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
}
