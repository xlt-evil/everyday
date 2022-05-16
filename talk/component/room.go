package component

/**
 * @Author: hxy
 * @Description:
 * @File:  room
 * @Date: 2022/04/07 10:01
 */

// 聊天室
type Room struct {
	RoomId int64
	Owner int64
	control SupervisorControl
	JoinQueue chan *ChatConn
}

func (r *Room) joinRoom() {
	for v := range r.JoinQueue {
		r.control.Add(v)
		v.Room = r // 互相知道对方关联
	}
}

func (r *Room) LeaveRoom() {

}