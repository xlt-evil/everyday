package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  proxy
 * @Date: 2022/04/19 13:47
 */

type ProxyHttp struct {
	Host string
	Port string
}


func (t *ProxyHttp) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Println("获取到一个请求")
	remote,err := url.Parse(fmt.Sprintf("http://%s:%s",t.Host,t.Port))
	if err != nil {
		panic(err)
	}
	ts := &http.Transport{
		Proxy: http.ProxyURL(remote),
	}
	c := http.Client{Transport: ts}
	r.RequestURI = "www.baidu.com"
	ws ,err := c.Do(r)
	if err != nil {
		w.Write([]byte("over"))
	}
	io.Copy(w,ws.Body)
}

func startServer() {
	proxy := &ProxyHttp{
		Host: "127.0.0.1",
		Port: "10809",
	}
	h := &http.Server{
		Addr:              ":20000",
		Handler:           proxy,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	//h := &ProxyHttp{
	//	Host: "127.0.0.1",
	//	Port: "10809",
	//}

	if err := h.ListenAndServe();err != nil {
		panic(err)
	}
	fmt.Println("代理服务器关闭")
}

//func Proxy() {
//	u,_ := url.Parse("http://127.0.0.1:10809")
//	t := &http.Transport{
//		Proxy: http.ProxyURL(u),
//	}
//	c := &http.Client{
//		Transport: t,
//		Timeout: time.Duration(10) * time.Second,
//	}
//	r,_ := c.Get("https://youtube.com/")
//	b,_ := ioutil.ReadAll(r.Body)
//	fmt.Println(string(b))
//	httpProxy := httputil.NewSingleHostReverseProxy(u)
//
//}

func main()  {
	startServer()
}