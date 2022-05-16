package _0220209

/**
 * @Author: hxy
 * @Description:
 * @File:  singleFlight
 * @Date: 2022/02/09 12:42
 */

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"testing"
	"time"
)

var g singleflight.Group
func T() {
	v,_,shard := g.Do("1", func() (interface{}, error) {
		v := 1
		return v,nil
	})
	fmt.Println(shard)
	fmt.Println(v)
}

func Test_Flight(t *testing.T) {
	for i := 0;i<10;i++ {
		go T()
	}
	c := time.NewTimer(time.Second*3)
	select {
	case <-c.C:
		
	}
}