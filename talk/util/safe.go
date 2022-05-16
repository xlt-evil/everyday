package util

import "fmt"

/**
 * @Author: hxy
 * @Description:
 * @File:  safe
 * @Date: 2022/04/07 10:19
 */


func SafeGoFn(fn func()) {
	defer func() {
		fmt.Println(fn)
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()
	fn()
}