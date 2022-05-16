package _0211129

/**
 * @Author: hxy
 * @Description:
 * @File:  两数之和 II - 输入有序数组_test
 * @Date: 2021/11/29 14:52
 */

//@tag [双指针]
func twoSum(numbers []int, target int) []int {
	low, high := 0,len(numbers) - 1
	for low < high {
		sum := numbers[low] + numbers[high]
		if  sum == target{
			return []int{low,high}
		}else if sum > target {
			high--
		}else {
			low++
		}
	}
	return nil
}