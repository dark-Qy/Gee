package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// NewRouter 初始化
func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 将pattern转换为parts
func parsePattern(pattern string) []string {
	ps := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range ps {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (router *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	parts := parsePattern(pattern)

	// 如果不存在method对应的根节点，则新建一个
	if _, ok := router.roots[method]; !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(pattern, parts, 0)
	router.handlers[key] = handler
}

// 查找路由，同时存储动态匹配的参数
func (router *Router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	if root, ok := router.roots[method]; ok { // 如果不存在method对应的根节点，则返回nil
		nodes := root.search(searchParts, 0)
		if nodes != nil {
			// 将parts中的动态参数存储到params中
			parts := parsePattern(nodes.pattern)
			for index, part := range parts {
				if part[0] == ':' {
					params[part[1:]] = searchParts[index]
				}
				if part[0] == '*' && len(part) > 1 {
					params[part[1:]] = strings.Join(searchParts[index:], "/")
					break
				}
			}
			return nodes, params
		}
	}
	return nil, nil
}

// ServeHTTP 实现http.Handler接口
func (router *Router) handle(c *Context) {
	// 获取路由并解析路由参数
	n, params := router.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		// 将路由对应的处理函数添加到中间件中
		c.handlers = append(c.handlers, router.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
