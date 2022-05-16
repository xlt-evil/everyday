package _0211115

/**
 * @Author: hxy
 * @Description:
 * @File:  矩阵置零_test
 * @Date: 2021/11/15 17:39
 */
//@tag [数组]
func setZeroes(matrix [][]int)  {
	m,n := len(matrix),len(matrix[0])
	rows := make([]bool,m)
	cols := make([]bool,n)
	for i,_ := range matrix {
		for j,_ := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i,_ := range matrix {
		for j,_ := range matrix[i] {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}


