package _0211201

/**
 * @Author: hxy
 * @Description:
 * @File:  连续字符_test
 * @Date: 2021/12/01 9:57
 */

//@tag [快慢指针]
func maxPower(s string) int {
	maxLength := 0
	slow, quick := 0,0
	for quick < len(s){
		if s[slow] != s[quick] {
			slow = quick
		}else {
			quick++
			maxLength = max(maxLength, quick - slow)
		}
	}
	return maxLength
}