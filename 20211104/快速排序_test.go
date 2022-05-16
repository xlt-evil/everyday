package _0211104

import (
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  快速排序_test
 * @Date: 2021/11/04 15:46
 */
//@tag [排序]
// 快速排序
// 随机从数字组取一个基准数，通常取最后一个，然后当前数组的左边全小于该数，右边全大于改数
// 然后用这个规则不断的分区，最后剩下一个数的时候就派好了
// 他和归并排序不一样的地方是快排是自顶向下，归并是自下向上，并且快排是原地排序
// 快排保证了每个分区都和当前的基准值有一个大小关系
// 归并是保证每个分区自己是有序的
// O(NLogN)  => O(N^2)
func quickSort(nums []int,start,end int) {
	if start >= end {
		return
	}
	p := partition(nums,start,end)
	quickSort(nums,start,p - 1)
	quickSort(nums,p + 1,end)
}

func partition(arr []int,start,end int) int {
	// 基准数
	pivot := arr[end]
	// 区间值
	i := start
	for j := start;j < end ;j++ {
		if arr[j] < pivot {
			arr[i],arr[j] = arr[j],arr[i]
			i++
		}
	}
	arr[i],arr[end] = arr[end],arr[i]
	return i
}


func Test_quickSort(t *testing.T){
	n := []int{3,5,4,9,7,2}
	quickSort(n,0,len(n) - 1)
	PrintIntN(n)
}
