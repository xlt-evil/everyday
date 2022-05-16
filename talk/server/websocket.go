package server

import (
	"everyDay/talk/component"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  websocket
 * @Date: 2022/04/07 15:19
 */

// socket 连接的基本配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header["Origin"]
		if len(origin) == 0 {
			return true
		}
		_, err := url.Parse(origin[0])
		if err != nil {
			return false
		}
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	component.CreateConn(conn, component.Center)
}