package _0211104

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  有效的完全平方数_test
 * @Date: 2021/11/04 9:25
 */
//@tag [递归]

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	return check(num / 2,num)
}

func check(n int,nums int) bool{
	if n < 1 {
		return false
	}
	if n * n == nums {
		return true
	}
	if (n / 2) * (n / 2) > nums {
		return check(n / 2,nums)
	}
	return check(n - 1,nums)
}

func Test_isPerfectSquare(t *testing.T) {
	fmt.Println(isPerfectSquare(1))
}

