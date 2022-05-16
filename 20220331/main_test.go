package _0220331

import (
	"fmt"
	"net"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  main
 * @Date: 2022/04/02 16:13
 */
func main()  {
	t,err := net.Dial("udp","192.168.5.246:554")
	if err != nil {
		panic(err)
	}
	b := make([]byte,1024)
	t.Read(b)
	fmt.Println(string(b))
}

func Test_main(t *testing.T) {
	main()
}