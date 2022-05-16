package _0211117

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  最大单词长度乘积_test
 * @Date: 2021/11/17 9:45
 */

//@tag [位运算]
func maxProduct(words []string) int {
	mask := make([]int32,len(words))
	for i,_ := range words {
		for _,v := range words[i] {
			mask[i] |= 1 <<  (v - 'a')
		}
	}
	max := 0
	for i,_ := range words {
		for j := i + 1;j < len(words);j++ {
			if mask[i] & mask[j] == 0 {
				max = Max(len(words[i]) * len(words[j]),max)
			}
		}
	}

	return max
}

func Max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_maxProduct(t *testing.T) {
	fmt.Println(119 & 126)
	fmt.Println(maxProduct([]string{"abcw","baz","foo","bar","xtfn","abcdef"}))
}