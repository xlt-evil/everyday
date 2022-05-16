package _0211116

import "sort"

/**
 * @Author: hxy
 * @Description:
 * @File:  有效的字母异位词_test
 * @Date: 2021/11/16 11:24
 */
func isAnagram(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}