package _0211103

import "testing"

/**
 * @Author: hxy
 * @Description:
 * @File:  归并排序_tesp
 * @Date: 2021/11/03 15:21
 */
//@tag [排序]
// 1.归并排序是将排序的数组分化，然后在合并分而治之
// 2.他不是原地的稳定排序
// 3.排序时间恒定，O(NLogN) 空间O(N)
// 合并的时候，需要先申请一个临时空间，然后左右区间从起始下标开始比较，小的那个存入临时空间，下标移动，直到某一区间结束，然后在把剩下空间一起并入
func MergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSort(arr,start,mid)
	MergeSort(arr,mid + 1,end)
	Merge(arr,start,mid,end)
}


func Merge(num []int,start,mid,end int) {
	t := []int{}
	i1,i2 := start,mid + 1
	for i1 <= mid && i2 <= end {
		if num[i1] < num[i2] {
			t = append(t,num[i1])
			i1 ++
		}else {
			t = append(t,num[i2])
			i2 ++
		}
	}
	if i1 <= mid {
		t = append(t,num[i1:mid + 1]...)
	}
	if i2 <= end {
		t = append(t,num[i2:end + 1]...)
	}
	for i,v := range t {
		num[start + i ] = v
	}
}

func Test_MergeSort(t *testing.T) {
	n := []int{9,8,7,6,3,4,2,1}
	MergeSort(n,0,len(n) - 1)
	PrintIntN(n)
}