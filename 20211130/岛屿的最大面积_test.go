package _0211130

/**
 * @Author: hxy
 * @Description:
 * @File:  岛屿的最大面积_test
 * @Date: 2021/11/30 10:30
 */

//@tag [深度优先搜索]
func maxAreaOfIsland(grid [][]int) int {
	var (
		dfs func(i,j int) int
		maxLength int
	)

	visitor := map[[2]int]bool{}
	dfs = func(i,j int) int {
		n := 0
		mark(visitor,i,j)
		if grid[i][j] == 0 {
			return n
		}
		if i + 1 < len(grid)  && !check(visitor,i + 1,j) {
			n += dfs(i + 1,j)
		}
		if i - 1 >= 0  && !check(visitor,i - 1,j) {
			n += dfs(i - 1,j)
		}
		if j + 1 < len(grid[i]) && !check(visitor,i,j + 1) {
			n += dfs(i,j + 1)
		}
		if j - 1 >= 0   && !check(visitor,i,j - 1) {
			n += dfs(i,j - 1)
		}
		return n + 1
	}
	for i,_ := range grid{
		for j,_ := range grid[i] {
			if check(visitor,i,j) {
				continue
			}
			maxLength = max(maxLength,dfs(i,j))
		}
	}
	return maxLength
}