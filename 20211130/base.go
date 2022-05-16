package _0211130

/**
 * @Author: hxy
 * @Description:
 * @File:  base
 * @Date: 2021/11/30 11:42
 */

func check(m map[[2]int]bool,i,j int) bool {
	return m[[2]int{i,j}]
}

func mark(m map[[2]int]bool,i,j int) {
	m[[2]int{i,j}] = true
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}