package model

/**
 * @Author: hxy
 * @Description:
 * @File:  req
 * @Date: 2022/04/07 14:10
 */

type JoinRoomMsg struct {
	RoomId int64 `json:"room_id"`// 房间号
}