package component

import (
	"encoding/json"
	"everyDay/talk/model"
	"everyDay/talk/util"
	"github.com/gorilla/websocket"
	"log"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  chat_conn
 * @Date: 2022/04/07 10:11
 */


type ChatConn struct {
	Conn *websocket.Conn
	ReadMsg chan model.Msg
	WriteMsg chan model.Msg
	BroadcastMsg chan model.Msg
	Room *Room
	RoomCenter *RoomCenter
}

func CreateConn(conn *websocket.Conn,center *RoomCenter) *ChatConn {
	c := &ChatConn{
		Conn: conn,
		ReadMsg: make(chan model.Msg,100),
		WriteMsg: make(chan model.Msg,100),
		RoomCenter: center,
	}
	go util.SafeGoFn(c.read)
	go util.SafeGoFn(c.write)
	return c
}



func (c *ChatConn) SetWrite(m model.Msg){
	c.WriteMsg <-m
}

func (c *ChatConn) Close() {
	c.Conn.Close()
}

// 具体业务处理
func (c *ChatConn) job() {
	for {
		select {
		case v := <-c.ReadMsg:
			switch v.Tag {

			}
		}
	}
}

// 读数据
func (c *ChatConn) read() {
	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		m := model.Msg{}
		if err = json.Unmarshal(p,&m);err != nil {
			log.Println("不是约定的信息包")
			continue
		}
		// 还没加入房间，要加入房间
		if c.Room == nil{
			c.init(m)
		}else {
			// 业务处理
			c.ReadMsg<-m
		}
	}
}

// 写数据
func (c *ChatConn) write() {
	for v := range c.WriteMsg {
		if err := c.Conn.WriteJSON(v);err != nil {
			return
		}
	}
}

// 输入数据
func (c *ChatConn) init(msg model.Msg) {
	if msg.Tag == model.Join {
		req := model.JoinRoomMsg{}
		b,_ := json.Marshal(msg.Data)
		_ = json.Unmarshal(b,&req)
		room,ok := c.RoomCenter.FindRoom(req.RoomId)
		if !ok {
			c.SetWrite(model.Msg{
				Tag: model.BusinessError,
				Err:    model.BusinessErr{
					Code: 100,
					Msg:  "房间号没有找到",
				},
			})
			return
		}
		room.JoinQueue <- c
	}
}

