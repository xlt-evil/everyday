package _0211207

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  边界着色_test
 * @Date: 2021/12/07 15:31
 */

//@tag [广度优先遍历]
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	visitor := make(map[[2]int]bool)
	queue := [][2]int{{row,col}}
	m,n := len(grid),len(grid[0])
	tag := [][2]int{}
	c := grid[row][col]
	for len(queue) != 0 {
		row = queue[0][0]
		col = queue[0][1]
		queue = queue[1:]
		mark(visitor,row,col)
		flag := false
		for _,v := range dirs {
			x,y := row + v[0],col + v[1]
			if x >= 0 && x < m && y >= 0 && y < n && grid [x][y] == c{
				if !check(visitor,x,y) {
					queue = append(queue,[2]int{x,y})
				}
			}else {
				flag = true
			}
			if flag {
				tag = append(tag,[2]int{row,col})
			}
		}
	}
	for _,v := range tag {
		grid[v[0]][v[1]] = color
	}
	return grid
}

func Test_colorBorder(t *testing.T) {
	fmt.Println(colorBorder([][]int{{1,2,1,2,1,2},{2,2,2,2,1,2},{1,2,2,2,1,2}},1,3,1))
}


//[1,2,1,2,1,2]
//[2,2,2,2,1,2]
//[1,2,2,2,1,2]
//
//[1,1,1,1,1,2]
//[1,1,1,1,1,2]
//[1,1,1,1,1,2]
//
//[1,1,1,1,1,2]
//[1,2,1,1,1,2]
//[1,1,1,1,1,2]