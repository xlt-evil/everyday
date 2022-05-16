package _0211115

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  杨辉三角_test
 * @Date: 2021/11/15 15:52
 */

//@tag [数组]
func generate(numRows int) [][]int {
	answer := make([][]int,numRows)
	for i := range answer {
		t := make([]int,i + 1)
		t[0],t[i] = 1,1
		for j := 1;j < i ;j ++ {
			t[j] = answer[i - 1][j - 1] + answer[i - 1][j]
		}
		answer[i] = t
	}
	return answer
}

func Test_generate(t *testing.T) {
	fmt.Println(generate(10))
}