package main

import (
	"Gee/gee"
	"fmt"
	"net/http"
)

func main() {
	// 首先对路由进行初始化
	r := gee.Default()

	// 然后根据不同的路由类型，调用不同的处理函数
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	// 启动路由
	r.Run(":8080")
}
