package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  计算某字符出现次数
 * @Date: 2022/05/16 11:36
 */

func CountChar(msg string, find string) (count int) {
	msg = strings.ToLower(msg)
	find = strings.ToLower(find)
	length := len(find)
	for i := 0; i < len(msg); {
		if i+length > len(msg) {
			return
		}
		if find == msg[i:i+length] {
			count++
			i = i + length
		} else {
			i++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	questionLine, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	answerLine, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(CountChar(strings.TrimSpace(questionLine), strings.TrimSpace(answerLine)))
}
