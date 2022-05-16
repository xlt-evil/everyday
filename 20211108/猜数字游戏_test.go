package _0211108

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  猜数字游戏_test
 * @Date: 2021/11/08 9:48
 */
//@tag[哈希]
func getHint(secret string, guess string) string {
	m1,m2 := make(map[uint8]int),make(map[uint8]int)
	A,B := 0,0
	for i,_ := range secret {
		if secret[i] == guess[i] {
			A++
		}
		m1[secret[i]]++
		m2[guess[i]]++
	}
	for k, v := range m1 {
		B += min(m2[k],v)
	}
	return fmt.Sprintf("%dA%dB",A,B - A)
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func Test_getHint(t *testing.T) {
	fmt.Println(getHint("1807","1807"))
}