package _0220329

/**
 * @Author: hxy
 * @Description:
 * @File:  旋转数组的最小数字_test
 * @Date: 2022/03/29 10:39
 */

// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
func minArray(numbers []int) int {
	for i := 0 ; i < len(numbers) - 1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}