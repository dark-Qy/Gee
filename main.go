package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 创建两个路由，第一个路由直接打印当前的路径，第二个打印当前路由的请求头所有内容
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// 启动监听程序
	http.ListenAndServe(":8080", nil)
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Method = %q\n", req.Method)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	// 循环打印req.Header中的所有内容
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}
