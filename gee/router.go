package gee

import (
	"fmt"
)

type Router struct {
	handles map[string]HandlerFunc
}

// 初始化
func NewRouter() *Router {
	return &Router{handles: make(map[string]HandlerFunc)}
}

func (router *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	router.handles[key] = handler
}

// ServeHTTP 实现http.Handler接口
func (router *Router) handle(c *Context) {
	key := c.Req.Method + "-" + c.Req.URL.Path
	// 根据key查找是否存在该路由
	if handler, ok := router.handles[key]; ok {
		handler(c.Writer, c.Req)
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %s\n", c.Req.URL)
	}
}
