package _0211129

/**
 * @Author: hxy
 * @Description:
 * @File:  无重复字符的最长子串_test
 * @Date: 2021/11/29 17:31
 */

//@tag [滑动窗口]
func lengthOfLongestSubstring(s string) int {
	var maxLength int
	m := make(map[uint8]int,len(s))
	r := -1
	for l,_ := range s {
		for r + 1 < len(s) && m[s[r + 1]] == 0 {
			m[s[r+1]]++
			r++
		}
		maxLength = max(len(m) ,maxLength)
		delete(m,s[l])
	}
	return maxLength
}

func max (a,b int)int {
	if a > b {
		return a
	}
	return b
}