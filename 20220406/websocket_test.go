package _0220406

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  websocket_test
 * @Date: 2022/04/06 13:56
 */

type Queue struct {
	conn map[string]*websocket.Conn
}

var Q Queue

func init() {
	Q = Queue{conn: map[string]*websocket.Conn{}}
}

func (c *Queue) Add(conn *websocket.Conn) {
	c.conn[conn.RemoteAddr().String()] = conn
}

func (c *Queue) Remove(conn *websocket.Conn) {
	delete(c.conn,conn.RemoteAddr().String())
	fmt.Println(conn.RemoteAddr().String() + "离线")
}

func (c *Queue) Length() int {
	return len(c.conn)
}

func (c *Queue) Broadcast(self string,msg []byte) {
	for _,v := range c.conn {
		self = strings.Split(self,":")[0]
		msg := []byte(self +"说:"+ string(msg))
		if err := v.WriteMessage(websocket.TextMessage,msg);err != nil {
			fmt.Println(fmt.Sprintf("远程连接 %s,发送消失失败 %s",v.RemoteAddr().String(),err))
		}
	}
}

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

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	sign := make(chan *websocket.Conn,5)
	Q.Add(conn)
	go read(conn,sign)
	go close(sign)
}

func read(conn *websocket.Conn,sign chan *websocket.Conn) {
	defer func() {
		sign <- conn
	}()
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		Q.Broadcast(conn.RemoteAddr().String(),p)
	}
}



func close(sign chan *websocket.Conn) {
	select {
	case c:=<-sign:
		Q.Remove(c)
	}
}

func Test_websocket(t *testing.T) {
	http.HandleFunc("/ws",handler)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println(err)
	}
}