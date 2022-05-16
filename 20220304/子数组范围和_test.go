package _0220304

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  子数组范围和_test
 * @Date: 2022/03/04 10:37
 */

func subArrayRanges(nums []int) int64 {
	m := map[string]string{}
	for i := 0 ;i <10;i++ {
		go func(i int) {
			m[strconv.Itoa(i)] = strconv.Itoa(i)
		}(i)
	}
	time.Sleep(time.Second * 10)
	fmt.Println("123")
	return 0
}

func Test_subArrayRanges(t *testing.T) {
	subArrayRanges(nil)
}