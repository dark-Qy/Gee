package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 中间件
	parent      *RouterGroup  // 支持嵌套
	engine      *Engine       // 所有的group共享一个Engine实例
}

// Engine 创建路由引擎
type Engine struct {
	*RouterGroup // 继承RouterGroup的属性
	router       *Router
	groups       []*RouterGroup
}

// Default 创建路由引擎
func Default() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 创建路由组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 添加路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	// 不同组需要进行区分
	pattern := group.prefix + comp
	// 调用router接口的addRoute方法
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 添加GET请求路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 添加POST请求路由
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
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
