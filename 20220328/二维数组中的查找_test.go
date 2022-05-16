package _0220328

/**
 * @Author: hxy
 * @Description:
 * @File:  二维数组中的查找_test
 * @Date: 2022/03/28 17:43
 */

// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	row,col := 0,cols - 1
	for ;row < rows && col >= 0; {
		e := matrix[row][col]
		if e == target {
			return true
		}else if e > target {
			col--
		}else {
			row++
		}
	}
	return false
}