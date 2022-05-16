package _0211209

import (
	"fmt"
	"strings"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  有效的井字游戏_test
 * @Date: 2021/12/09 10:00
 */

func win(board []string, p byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == p && board[i][1] == p && board[i][2] == p ||
			board[0][i] == p && board[1][i] == p && board[2][i] == p {
			return true
		}
	}
	return board[0][0] == p && board[1][1] == p && board[2][2] == p ||
		board[0][2] == p && board[1][1] == p && board[2][0] == p
}

func validTicTacToe(board []string) bool {
	oCount, xCount := 0, 0
	for _, row := range board {
		oCount += strings.Count(row, "O")
		xCount += strings.Count(row, "X")
	}

	return !(oCount != xCount && oCount != xCount-1 || // 合法的棋在数量上总是等于 或则 X多一个
		     oCount != xCount && win(board, 'O') || // 如果O赢了，则说明 o 一定要等于 x
		     oCount != xCount-1 && win(board, 'X')) // 如果 x 赢了，则说明x 一定多一个
}


func Test_validTicTacToe(t *testing.T) {
	fmt.Println(validTicTacToe([]string{"XOX","O O","XOX"}))
}