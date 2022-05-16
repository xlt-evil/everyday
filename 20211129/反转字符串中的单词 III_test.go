package _0211129

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  反转字符串中的单词 III
 * @Date: 2021/11/29 16:49
 */


//@tag 双指针
func reverseWords(s string) string {
	ans := make([]byte,len(s))
	j := 0
	for i := 0;i < len(s);i++ {
		low,high := -1,-1
		if s[i] == ' '  {
			low ,high = j,i - 1
			ans[i] = s[i]
			j = i + 1
		}else if i == len(s) - 1 {
			low ,high = j,i
			j = i + 1
		}else {
			continue
		}
		for low <= high {
			ans[low],ans[high] = s[high],s[low]
			low++
			high--
		}
	}
	return string(ans)
}

func Test_reverseWord(t *testing.T) {
	fmt.Println(reverseWords("Let's take LeetCode contest") )
	fmt.Println("s'teL ekat edoCteeL tsetnoc")
}