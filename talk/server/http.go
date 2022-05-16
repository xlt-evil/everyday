package server

import "net/http"

/**
 * @Author: hxy
 * @Description:
 * @File:  http
 * @Date: 2022/04/07 15:20
 */

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("roomId")
}