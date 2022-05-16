package _0220307

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  七进制数_test
 * @Date: 2022/03/07 10:57
 */

//链接：https://leetcode-cn.com/problems/base-7/solution/qi-jin-zhi-shu-by-leetcode-solution-cl2v/
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negative := num < 0
	if negative {
		num = -num
	}
	s := []byte{}
	for num > 0 {
		s = append(s, '0'+byte(num%7))
		num /= 7
	}
	if negative {
		s = append(s, '-')
	}
	for i, n := 0, len(s); i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return string(s)
}

func Test_covert(t *testing.T) {
	fmt.Println(convertToBase7(100))
}