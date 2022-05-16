package _0211115

import (
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  有效的数独_test
 * @Date: 2021/11/15 16:15
 */

//@tag [数组]
func isValidSudoku(board [][]byte) bool {
	var (
		rows [10][10]bool
		cols [10][10]bool
		area [10][10]bool
	)
	for i := 0;i < 9;i++ {
		for j := 0;j < 9;j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}
			u := c - '0'
			idx := i / 3 * 3 + j / 3
			if rows[i][u] || cols[j][u] || area[idx][u] {
				return false
			}
			rows[i][u],cols[j][u],area[idx][u] = true,true,true
		}
	}
	return true
}

func Test(t *testing.T) {
}