package _0211202

/**
 * @Author: hxy
 * @Description:
 * @File:  01 矩阵_test
 * @Date: 2021/12/02 9:28
 */

//@tag [广度优先搜索]
func updateMatrix(mat [][]int) [][]int {
	visitors := map[[2]int]bool{}
	queue := [][2]int{}
	ans := make([][]int,len(mat))
	for i,_ := range mat {
		ans[i] = make([]int,len(mat[i]))
		for j,_ := range mat[i] {
			if mat[i][j] == 0 {
				queue = append(queue,[2]int{i,j})
				mark(visitors,i,j)
			}
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0;i < 4;i++ {
			dx := cur[0] + dirs[i][0]
			dy := cur[1] + dirs[i][1]
			if dx >= 0 && dx < len(mat) && dy >= 0 && dy < len(mat[cur[0]]) && !check(visitors,dx,dy){
				ans[dx][dy] = ans[cur[0]][cur[1]] + 1
				queue = append(queue,[2]int{dx,dy})
				mark(visitors,dx,dy)
			}
		}
	}

	return ans
}


