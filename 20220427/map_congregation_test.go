package _0220427

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  map_congregation
 * @Date: 2022/04/27 11:34
 */




func Test(t *testing.T)  {
	var (
		m1 = map[int]int{1:1,2:2,3:3}
		m2 = map[int]int{1:1,2:3,3:5}
		diff  []int // 差
		same  []int // 交
	)
	for k,v:= range m1{
		y,ok := m2[k]
		if !ok || v !=y {
			diff = append(diff, v)
		}else {
			same = append(same, v)
		}
	}
	fmt.Println("m1 和 m2 的交集",same)
	for k,_:= range m2{
		if _,ok:=m1[k];!ok {
			diff = append(diff,k)
		}
	}
	fmt.Println("m1 和m2的差集",diff)
}