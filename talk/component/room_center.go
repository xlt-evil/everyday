package component

import (
	"everyDay/talk/model"
	"github.com/gorilla/websocket"
	"sync"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  room_center
 * @Date: 2022/04/07 10:41
 */

// 聊天室中心
type RoomCenter struct {
	Rooms sync.Map
	TempConn sync.Map // 临时大厅连接中心
}

var Center *RoomCenter

func init() {
	Center = &RoomCenter{
		Rooms:    sync.Map{},
		TempConn: sync.Map{},
	}
	r := &Room{
		RoomId:    1,
		Owner:     0,
		control:   SupervisorControl{BroadcastMsg: make(chan model.Msg),},
		JoinQueue: make(chan *ChatConn,100),
	}
	go r.joinRoom()
	Center.AddRoom(r)
}

func (r *RoomCenter) FindRoom(roomId int64) (*Room,bool) {
	room,ok := r.Rooms.Load(roomId)
	if !ok {
		return nil,false
	}
	v := room.(*Room)
	return v,ok
}

func (r *RoomCenter) AddRoom(room *Room) {
	r.Rooms.Store(room.RoomId,room)
}

func (r *RoomCenter) Delete(roomId int64) {
	r.Rooms.Delete(roomId)
}


func (r *RoomCenter) AddConn(chat *websocket.Conn) {
	r.TempConn.Store(chat.RemoteAddr().String(),chat)
}

