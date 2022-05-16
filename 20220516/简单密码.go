package main

import "fmt"

/**
 * @Author: hxy
 * @Description:
 * @File:  简单密码
 * @Date: 2022/05/16 16:19
 */

const lower = 'a' - 'A'

var dist = map[rune]int32{
	'a': 2, 'b': 2, 'c': 2,
	'd': 3, 'e': 3, 'f': 3,
	'g': 4, 'h': 4, 'i': 4,
	'j': 5, 'k': 5, 'l': 5,
	'm': 6, 'n': 6, 'o': 6,
	'p': 7, 'q': 7, 'r': 7, 's': 7,
	't': 8, 'u': 8, 'v': 8,
	'w': 9, 'x': 9, 'y': 9, 'z': 9,
}

func getWordType(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r = r + lower + 1
		if r > 'z' {
			r = 'a'
		}
	} else if r >= 'a' && r <= 'z' {
		r = dist[r] + '0'
	}
	return r
}

func Encrypt(pwd string) string {
	list := []rune(pwd)
	for i, _ := range list {
		list[i] = getWordType(list[i])
	}
	return string(list)
}

func main() {

	var pwd string
	if _, err := fmt.Scanf("%s", &pwd); err != nil {
		panic(err)
	}
	fmt.Println(Encrypt(pwd))
}
