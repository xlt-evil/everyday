package _0211202

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  腐烂的橘子_test
 * @Date: 2021/12/02 11:02
 */

//@tag [广度优先搜索]
func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	time := 0
	sum := 0
	visitors := map[[2]int]bool{}
	queue := make([][2]int,0)
	for i,_ := range grid {
		for j,_ := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue,[2]int{i,j})
				mark(visitors,i,j)
			}else if grid[i][j] == 1 {
				sum ++
			}
		}
	}

	sum1 := 0
	for len(queue) > 0 {
		length := len(queue)
		flag := false
		for i := 0;i < length;i++ {
			for d := 0 ;d < 4;d++ {
				dx := queue[i][0] + dirs[d][0]
				dy := queue[i][1] + dirs[d][1]
				if dx >= 0 && dx < m && dy >= 0 && dy < n && !check(visitors,dx,dy) && grid[dx][dy] == 1{
					flag = true
					mark(visitors,dx,dy)
					queue = append(queue,[2]int{dx,dy})
					sum1 ++
				}
			}
		}
		if flag {
			time++
		}
		queue = queue[length:]
	}
	if sum != sum1 {
		time = -1
	}
	return time
}

func Test_orangesRotting(t *testing.T) {
	fmt.Println(orangesRotting([][]int{{2,1,1},{1,1,0},{0,1,1}}))
}