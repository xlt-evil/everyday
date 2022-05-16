package _0220225

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  context_test
 * @Date: 2022/02/25 14:21
 */


func ContextCancel() {
	ctx,cancel := context.WithCancel(context.TODO())
	go Demo(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second)
}

func Demo(ctx context.Context) {
	ticker := time.NewTimer(time.Second * 10)

	select {
	case <-ticker.C:
		fmt.Println("timer out")
	//case <-ctx.Done():
	//	fmt.Println("context is cancel")
	}
}

func Test_context(test *testing.T) {
	ContextCancel()
}