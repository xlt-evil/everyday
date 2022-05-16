package component

import (
	"everyDay/talk/model"
	"sync"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  queue
 * @Date: 2022/04/07 9:58
 */

// 连接管理中心
type SupervisorControl struct {
	length int64
	conn sync.Map
	BroadcastMsg  chan model.Msg
}

// 添加进入房间人数
func (c *SupervisorControl) Add(conn *ChatConn) {
	defer func() {c.length++}()
	c.conn.Store(conn.Conn.RemoteAddr().String(),conn)
	conn.BroadcastMsg = c.BroadcastMsg // 建立连接通道
}

// 移除下线连接
func (c *SupervisorControl) Remove(conn *ChatConn) {
	defer func() {c.length--}()
	c.conn.Delete(conn.Conn.RemoteAddr().String())
}

// 当前连接数
func (c *SupervisorControl) Length() int64 {
	return c.length
}

// 广播
func (c *SupervisorControl) Broadcast() {
	for v := range c.BroadcastMsg{
		c.conn.Range(func(key, value interface{}) bool {
			x := value.(*ChatConn)
			x.SetWrite(v)
			return true
		})
	}
}