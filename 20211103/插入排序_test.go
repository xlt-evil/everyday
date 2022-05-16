package _0211103

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  插入排序_test
 * @Date: 2021/11/03 13:52
 */
//@tag [排序]
// 左右区间 ，左边是排序好的，右边是未排序的
// 默认1是已经排序好的
// 每次从未排序的拿一个出来，到已排序里一一比对找到正确的位置
func InsertionSort(nums []int) {
	for i := 1;i < len(nums);i++ {
		// 准备排序的元素
		value := nums[i]
		j := i - 1 // 已排序的区间
		// 默认当前元素应该是最大的，只和前一个比较
		for ; j >= 0 ;j -- {
			// 不满足，所以当前元素需要插入到之前已经排序中
			if nums[j] > value {
				nums[j + 1] = nums[j]
			}else {
				break
			}
		}
		// 前面已经排过序列,所以j + 1 是当前应该插入的下标
		nums[j + 1] = value
	}
}

func Test_InsertionSort(t *testing.T) {
	n := []int{4,6,3,2,1}
	InsertionSort(n)
	PrintIntN(n)
}