package _0220307

import (
	"fmt"
	"net/http"
	"strings"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  proxy_test
 * @Date: 2022/03/08 17:39
 */

import (
	"io/ioutil"
	"log"
	"net/url"
	"testing"
)

var proxyConf = "127.0.0.1:19180"

func buildHtppClient(isProxy bool) *http.Client {
	var proxy func(*http.Request) (*url.URL, error) = nil
	if isProxy {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + proxyConf)
		}
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	return client
}

func Test_crawler(t *testing.T) {
	MyHttp()
}

func MyHttp() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//url := "https://www.google.com/"
		fmt.Println(request.URL.Path)
		client := buildHtppClient(true)
		if strings.Contains(request.URL.String(),"google") {
			println("okokokoko")
		}
		req, err := http.NewRequest("GET", request.URL.String(), nil)
		if err != nil {
			log.Println(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}


		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		}

		writer.Write(body)
	})
	http.ListenAndServe("192.168.3.63:9090",nil)
}