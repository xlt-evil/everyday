package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  字符串排序_test
 * @Date: 2022/05/16 14:51
 */

func StringSort(list []string) {
	sort.Strings(list)
}

func main() {
	n := 0
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	var list []string
	for i := 0; i < n; i++ {
		var s string
		if _, err = fmt.Scanf("%s", &s); err != nil {
			panic(err)
		}
		list = append(list, s)
	}
	StringSort(list)
	for _, v := range list {
		fmt.Println(v)
	}
}
