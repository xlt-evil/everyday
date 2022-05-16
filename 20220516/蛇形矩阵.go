package main

import (
	"fmt"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  蛇形矩阵
 * @Date: 2022/05/16 10:59
 */

// https://www.nowcoder.com/practice/649b210ef44446e3b1cd1be6fa4cab5e?tpId=37&t
func snakeSquare(n int) (result [][]int) {
	var (
		start = 1
		r = 0
		c = 2
	)
	for i := 0 ;i < n ; i ++ {
		rows := make([]int,0)
		start = start + r
		rows = append(rows,start)
		c1 := c
		for j := 1 ;j < n - i ; j++ {
			rows = append(rows,rows[j - 1 ] + c1)
			c1++
		}
		r++
		c++
		result = append(result,rows)
	}
	return result
}

func MyPrint(result [][]int) {
	for _,rows := range result {
		for _,cols := range rows {
			fmt.Print(cols," ")
		}
		fmt.Print("\n")
	}
}


func main() {
	n := 0
	if _,err := fmt.Scanf("%d",&n);err != nil {
		panic(err)
	}
	MyPrint(snakeSquare(n))
}