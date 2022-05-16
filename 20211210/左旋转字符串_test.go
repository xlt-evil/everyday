package _0211210

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  左旋转字符串_test
 * @Date: 2021/12/10 11:44
 */


//@tag [字符串]
func reverseLeftWords(s string, n int) string {
	k := n % len(s)
	return s[k:] + s[:k]
}

func Test_reverseLeftWords(t *testing.T) {
	fmt.Println(reverseLeftWords("abcdefg",10))
}