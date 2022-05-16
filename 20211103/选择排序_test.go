package _0211103

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  选择排序_test
 * @Date: 2021/11/03 14:58
 */
//@tag [排序]
// 1. 选择排序也分左右已排序和未排序区间
// 2. 每次都从未排序区间找到一个最小的和之前的交换位置
// 3. 他是原地的是稳定排序
func SelectionSort(nums []int) {
	for i := 0 ;i < len(nums);i++ {
		min := i
		for j := i + 1; j < len(nums) ; j ++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i],nums[min] = nums[min],nums[i]
	}
}


func Test_SelectionSort(t *testing.T) {
	n := []int{3,5,4,9,7,2}
	SelectionSort(n)
	PrintIntN(n)
}