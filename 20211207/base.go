package _0211207

/**
 * @Author: hxy
 * @Description:
 * @File:  base
 * @Date: 2021/12/07 15:33
 */

func check(m map[[2]int]bool,i,j int) bool {
	return m[[2]int{i,j}]
}

func mark(m map[[2]int]bool,i,j int) {
	m[[2]int{i,j}] = true
}

var dirs = [][2]int{{1,0},{-1,0},{0,1},{0,-1}}