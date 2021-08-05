package main

import (
	"fmt"
	"log"
	"net/http"
)

//定义空结构体
type Engine struct {
}

//实现ServeHTTP方法 有两个参数 一个是请求的响应  一个是请求信息 注意是 ServeHTTP 不是 ServerHTTP
func (engin *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path= %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Printf("test")
	}
}

func main() {
	engine := new(Engine)
	//使用自定义的http handler
	log.Fatal(http.ListenAndServe(":9999", engine))
}
