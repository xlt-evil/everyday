package _0220331

/**
 * @Author: hxy
 * @Description:
 * @File:  自除数_test
 * @Date: 2022/03/31 22:38
 */

// https://leetcode-cn.com/problems/self-dividing-numbers/
func selfDividingNumbers(left int, right int) []int {
	ans := []int{}
	for i := left;i <=right;i++ {
		if mod(i) {
			ans = append(ans,i)
		}
	}
	return ans
}

func mod(num int) bool {
	x := num
	for x != 0 {
		n := x % 10
		if n == 0 {
			return false
		}
		if num % n != 0 {
			return false
		}
		x = x / 10
	}
	return true
}