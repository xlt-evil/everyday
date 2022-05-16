package _0211125

/**
 * @Author: hxy
 * @Description:
 * @File:  轮转数组_test
 * @Date: 2021/11/25 13:51
 */
//@tag [数学]
func rotate(nums []int, k int)  {
	k = len(nums) - k % len(nums)
	var ans []int
	ans = append(ans,nums[k:]...)
	ans = append(ans,nums[:k]...)
	copy(nums,ans)
}