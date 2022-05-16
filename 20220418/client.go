package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  client
 * @Date: 2022/04/19 14:14
 */

func main() {
	url,_ := url2.Parse("http://127.0.0.1:20000")

	t := &http.Transport{
		Proxy: http.ProxyURL(url),
	}
	c := http.Client{Transport: t}
	b,err := c.Get("http://www.baidu.com/")
	if err != nil {
		return
	}
	s,err := ioutil.ReadAll(b.Body)
	if err != nil {
		return
	}
	fmt.Println(string(s))
}