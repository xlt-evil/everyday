package _0220113

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  至少是其他数字两倍的最大数_test
 * @Date: 2022/01/13 13:35
 */

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	max,secondMax,index := -1,-1,-1
	for i,v := range nums {
		if v > max {
			secondMax,max,index = max,v,i
		}else if v > secondMax {
			secondMax = v
		}
	}
	if max >= secondMax * 2 {
		return index
	}
	return -1
}

func Test_dominantIndex(t *testing.T) {
	fmt.Println(dominantIndex([]int{1,2,3,4}))
}