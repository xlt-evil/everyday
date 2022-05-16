package _0220419

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  crpyto
 * @Date: 2022/04/19 20:53
 */

var word = []rune{'A','B','C','D','E','F','G','H','I','J','K','L',}
func Crypto(name string) string {
	b := []rune(name)
	if len(name) % 2 == 0 {

	}
	for i,_ := range b {
		fmt.Println(b[i])
	}
	return ""
}



func Test_Crypto(t *testing.T) {
	fmt.Println(Crypto("糊糊安德森破卡跑得快"))
}