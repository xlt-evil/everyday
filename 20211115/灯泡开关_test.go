package _0211115

import (
	"fmt"
	"math"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  灯泡开关_test
 * @Date: 2021/11/15 9:44
 */

//@tag [数学]
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}

func Test_bulbSwitch(t *testing.T) {
	fmt.Println(bulbSwitch(3))
}