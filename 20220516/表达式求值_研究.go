/**
 * @Author: hxy
 * @Description:
 * @File:  表达式求值
 * @Date: 2022/05/16 17:14
 */

//给定一个字符串描述的算术表达式，计算出结果值。
//
//输入字符串长度不超过 100 ，合法的字符包括 ”+, -, *, /, (, )” ， ”0-9

// todo
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// todo 研究

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(OutPrint(input.Text()))
}

func OutPrint(s string) int {
	bytes := make([]byte, len(s))
	copy(bytes, s)
	return subCal(&bytes)
}

func subCal(bytes *[]byte) int {
	stack := []int{}
	num := 0
	var sigh byte = '+'
	for len(*bytes) > 0 {
		item := (*bytes)[0]
		*bytes = (*bytes)[1:]
		if unicode.IsDigit(rune(item)) {
			val, _ := strconv.Atoi(string(item))
			num = num*10 + val
		}
		if item == '(' {
			num = subCal(bytes)
		}
		if (!unicode.IsDigit(rune(item)) && item != ' ') || len(*bytes) == 0 {
			switch sigh {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				num *= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, num)
			case '/':
				num = stack[len(stack)-1] / num
				stack = stack[:len(stack)-1]
				stack = append(stack, num)
			}
			sigh = item
			num = 0
		}
		if item == ')' {
			break
		}
	}
	sum := 0
	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}
	return sum
}
