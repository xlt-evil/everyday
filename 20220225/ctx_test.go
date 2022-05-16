package _0220225

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  ctx_test
 * @Date: 2022/02/25 14:33
 */


func Ctx() {
	t := make(chan int)
	go func() {
		fmt.Println("123")

		close(t)
	}()
	select {
	case v,ok := <-t:
		fmt.Println(v,ok)
	}
}

func TestCtx(t *testing.T) {
	Ctx()
}