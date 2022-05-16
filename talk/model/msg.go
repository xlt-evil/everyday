package model

/**
 * @Author: hxy
 * @Description:
 * @File:  msg
 * @Date: 2022/04/07 9:55
 */

type Msg struct {
	Tag    MsgType `json:"tag"`
	Talker int64  `json:"talker"`
	Data   interface{} `json:"data"`
	Err    BusinessErr
}

type MsgType int

const (
	_             MsgType = iota
	Text                  // 文本
	Link                  // 链接
	Img                   // 图片
	Join                  // 加入房间
	Leave                 // 离开房间
	BusinessError         // 业务错误
)