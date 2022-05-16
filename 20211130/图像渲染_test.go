package _0211130

/**
 * @Author: hxy
 * @Description:
 * @File:  图像渲染_test
 * @Date: 2021/11/30 9:49
 */

//@tag [广度优先搜索]
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	visitor := map[[2]int]bool{}
	queue := [][2]int{{sr,sc}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		visitor[cur] = true
		i,j := cur[0],cur[1]
		// 上
		if i - 1 >= 0 && !check(visitor,i - 1,j) && image[i-1][j] == image[i][j]{
			queue = append(queue,[2]int{i - 1,j})
			mark(visitor,i - 1,j)
		}
		// 下
		if i + 1 <len(image) && !check(visitor,i + 1,j) && image[i + 1][j] == image[i][j]{
			queue = append(queue,[2]int{i + 1,j})
			mark(visitor,i + 1,j)
		}
		// 左
		if j - 1 >= 0 && !check(visitor,i,j - 1) && image[i][j - 1] == image[i][j] {
			queue = append(queue,[2]int{i,j - 1})
			mark(visitor,i,j - 1)
		}
		// 右
		if j + 1 < len(image[i]) && !check(visitor,i,j + 1) && image[i][j + 1] == image[i][j] {
			queue = append(queue,[2]int{i,j + 1})
			mark(visitor,i,j + 1)
		}
		image[i][j] = newColor
	}
	return image
}

