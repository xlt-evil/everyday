package _0220215

/**
 * @Author: hxy
 * @Description:
 * @File:  矩阵中的幸运数_test
 * @Date: 2022/02/15 11:44
 */

func luckyNumbers (matrix [][]int) []int {
	result := []int{}
	for i,_ := range matrix {
		min := 0
		flag := true
		for j,_ := range matrix[i] {
			if matrix[i][min] > matrix[i][j] {
				min = j
			}
		}
		for x := 0 ;x < len(matrix);x++ {
			if matrix[i][min] < matrix[x][min] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result,matrix[i][min])
		}
	}
	return result
}