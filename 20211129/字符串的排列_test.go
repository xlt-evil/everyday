package _0211129

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  字符串的排列_test
 * @Date: 2021/11/29 17:42
 */

func checkInclusion(s1 string, s2 string) bool {
	m1,m2 := make(map[uint8]int,0),make(map[uint8]int,0)
	for i,_ := range s1{
		m1[s1[i]]++
	}
	r := -1
	for i := 0 ;i < len(s2);i++ {
		for r + 1< len(s2) && r - i < len(s1) -1 {
			m2[s2[r + 1]]++
			r++
		}
		if equal(m1,m2) {
			return true
		}
		m2[s2[i]]--
	}
	return false
}

func equal(s1,s2 map[uint8]int) bool{
	for k,_ := range s1 {
		if s1[k] != s2[k] {
			return false
		}
	}
	return true
}

func Test_checkInclusion(t *testing.T) {
	fmt.Println(checkInclusion("adc","dcda"))
}