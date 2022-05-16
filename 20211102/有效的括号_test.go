package _0211102

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  有效的括号_test
 * @Date: 2021/11/02 14:00
 */
//@tag [栈]
func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	cur := []rune{}
	for _,v := range s{
		if v == '{' || v == '(' || v == '[' {
			cur = append(cur,v)
			continue
		}
		if len(cur) - 1  < 0 {
			return false
		}
		need := '0'
		switch v {
		case '}': need = '{'
		case ']': need = '['
		case ')': need = '('
		}
		t := cur[len(cur) - 1]
		cur = cur[:len(cur) - 1]
		if t != need {
			return false
		}
	}
	return len(cur) == 0
}
func Test_isValid(t *testing.T) {
	fmt.Println(isValid("((()))"))
	fmt.Println(isValid(")"))
	fmt.Println(isValid("[{()}]"))
	fmt.Println(isValid("()()(())"))

}