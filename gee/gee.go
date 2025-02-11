package gee

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 创建路由引擎
type Engine struct {
	router *Router
}

// Default 创建路由引擎
func Default() *Engine {
	return &Engine{router: NewRouter()}
}

// 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	// 调用router接口的addRoute方法
	engine.router.addRoute(method, pattern, handler)
}

// GET 添加GET请求路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加POST请求路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动路由引擎
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 实现http.Handler接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 创建上下文
	c := NewContext(w, r)
	engine.router.handle(c)
}
