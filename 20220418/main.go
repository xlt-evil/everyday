package main

import (
	"fmt"
	"net"
	"time"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  main
 * @Date: 2022/03/21 13:44
 */

func main() {
	sign := make(chan int)
	ip := "162.14.83.98"
	port := "520"
	u := fmt.Sprintf("%s:%s",ip,port)
	conn,err := net.DialTimeout("tcp",u,1 * time.Second)
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	go Read(conn,sign)
	go Write(conn)
	select {
	case <-sign:
		fmt.Println("over")
	}
}


func Read(conn net.Conn,close chan int) {
	i := 0
	for {
		b := make([]byte,1024 * 8 )
		n,err := conn.Read(b)
		if err != nil  {
			fmt.Println(err)
			i++
			if i > 3 {
				close <- 1
				break
			}
			continue
		}
		fmt.Println(string(b[:n]))
	}
}

func Write(conn net.Conn) {
	for {
		conn.Write([]byte(""))
		time.Sleep(time.Second * 10)
	}
}