package _0211129

/**
 * @Author: hxy
 * @Description:
 * @File:  反转字符串_test
 * @Date: 2021/11/29 16:35
 */

//@tag 双指针
func reverseString(s []byte)  {
	low,high := 0,len(s) - 1
	for low < high {
		s[low],s[high] = s[high],s[low]
		low++
		high--
	}
}