package server

import (
	"fmt"
	"net/http"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  server
 * @Date: 2022/04/07 14:27
 */



func Start() {
	fmt.Println("服务器监听在8888")
	http.HandleFunc("/ws",wsHandler)
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic("socket 服务启动失败！")
	}
}