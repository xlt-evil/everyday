package _0220113

import (
	"fmt"
	"math"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  整数反转_test
 * @Date: 2022/01/13 15:46
 */


//@tag [数学]
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}


func Test_reverse(t *testing.T) {
	fmt.Println(reverse(1534236469))
}