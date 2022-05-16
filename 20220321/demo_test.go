package _0220321

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  demo_test
 * @Date: 2022/03/21 13:40
 */

func Test_Demo(t *testing.T) {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for {
			i := <-abc
			fmt.Println(i)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(1 * time.Second)
}