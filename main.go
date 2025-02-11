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
	r.GET("/", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Req.URL.Path)
	})

	r.GET("/hello", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	// 启动路由
	r.Run(":8080")
}
