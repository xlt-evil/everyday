package _0211125

/**
 * @Author: hxy
 * @Description:
 * @File:  有序数组的平方_test
 * @Date: 2021/11/25 13:40
 */

//@tag [双指针]
func sortedSquares(nums []int) []int {
	i,j := 0,len(nums) - 1
	ans := make([]int,len(nums))
	for pos := len(nums) - 1;pos >= 0 ;pos-- {
		if v,w := nums[i] * nums[i],nums[j]*nums[j];v > w {
			ans[pos] = v
			i++
		}else {
			ans[pos] = w
			j--
		}
	}
	return ans
}