package _0211103

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  冒泡排序_test
 * @Date: 2021/11/03 9:47
 */

//@tag [排序]
// 原地排序稳定
func bubbleSort(nums [] int) {
	for i := 0 ;i < len(nums) ; i++ {
		// 每次冒一个
		for j := 0;j < len(nums) - i - 1;j ++ {
			if nums[j] > nums[j + 1] {
				nums[j],nums[j + 1] = nums[j + 1],nums[j]
			}
		}
	}
}

func Test_bubbleSort(t *testing.T){
	sort := []int{2,0,7,3,5,5,8}
	bubbleSort(sort)
	PrintIntN(sort)
}
