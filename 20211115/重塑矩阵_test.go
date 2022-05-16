package _0211115

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  重塑矩阵_test
 * @Date: 2021/11/15 14:50
 */

//@tag [数组]
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m,n := len(mat),len(mat[0])
	if m *n != r *c {
		return mat
	}
	sum := 0
	answer := make([][]int,r)
	for i := 0;i < r;i++ {
		t := make([]int,c)
		for j := 0;j < c;j++ {
			t[j] = mat[sum / n][sum % n]
			sum ++
		}
		answer[i] = t
	}
	return answer
}

func Test_matrixReshape(t *testing.T) {
	fmt.Println(matrixReshape([][]int{{1,2,3,4},{5,6,7,8}},4,2))
}