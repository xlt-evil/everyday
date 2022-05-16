package _0220209

/**
 * @Author: hxy
 * @Description:
 * @File:  打印_test
 * @Date: 2022/02/09 11:14
 */

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

type Instance struct {
	A string
	B int
	C *Inner
}

type Inner struct {
	D string
	C string
}

func Test_print(t *testing.T) {
	ins := Instance{
		A: "123",
		B: 1000,
		C: &Inner{
			D: "123",
			C: "456",
		},
	}
	spew.Dump(ins)
}