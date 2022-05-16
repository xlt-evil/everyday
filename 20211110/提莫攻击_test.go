package _0211110

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  提莫攻击_test
 * @Date: 2021/11/10 11:03
 */

func findPoisonedDuration(timeSeries []int, duration int) int {
	dp := -1
	sum := 0
	for _ ,v := range  timeSeries {
		if dp >= v {
			sum += v + duration - 1 - dp
		}else {
			sum += duration
		}
		dp = v + duration - 1
	}
	return sum
}

func Test_findPoisonedDuration(t *testing.T) {
	fmt.Println(findPoisonedDuration([]int{1,2},2))
}