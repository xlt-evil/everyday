package _0220211



/**
 * @Author: hxy
 * @Description:
 * @File:  pprof_test
 * @Date: 2022/02/11 17:10
 */

import (
	"net/http"
	_ "net/http/pprof"
	"testing"
	//"github.com/EDDYCJY/go-pprof-example/data"
)

func Test_net_pprof(t *testing.T) {
	go func() {
		for {
			select {

			}
			//time.Sleep(10 * time.Second)
			//go func() {
			//	t := time.NewTimer(10 * time.Second)
			//	select {
			//	case <-t.C:
			//
			//	}
			//}()
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}