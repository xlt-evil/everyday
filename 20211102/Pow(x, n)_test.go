package _0211102

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  Pow(x, n)_test
 * @Date: 2021/11/02 15:53
 */

//@tag [递归]

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n % 2 != 0 {
		return x * myPow(x, n - 1)
	}
	return myPow(x * x, n / 2)
}



func Test_myPow(t *testing.T) {
	fmt.Println(myPow(2,6))

}