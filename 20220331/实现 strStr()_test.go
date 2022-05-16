package _0220331

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  实现 strStr()_test
 * @Date: 2022/03/31 22:40
 */

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	l1 := len(haystack)
	l2 := len(needle)
	if l1 < l2 {
		return -1
	}
	for i := 0 ;i < len(haystack);i++ {
		if i + l2 > len(haystack) {
			return -1
		}
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}

func Test_strStr(t *testing.T) {
	fmt.Println(strStr("abc","c"))
}