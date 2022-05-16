package _0220303

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  各位数相加_test
 * @Date: 2022/03/03 17:44
 */


func addDigits(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
		if num == 0 && sum >= 10 {
			num,sum = sum,0
		}
	}
	return sum
}

func Test_addDigits(t *testing.T) {
	fmt.Println(addDigits(38))
}