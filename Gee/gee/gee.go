// engine结构实现了servehttp方法，get post 添加路由，运行，servehttp，本质上是保存路由表
package gee

import (
	"net/http"
	"strings"
)

// router map["method-url"] = handlerfunc
type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix   string
	engine   *Engine
	middlers []HandlerFunc
}

type HandlerFunc func(c *Context)

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix:   group.prefix + prefix,
		engine:   group.engine,
		middlers: make([]HandlerFunc, 0),
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Use(h ...HandlerFunc) {
	group.middlers = append(group.middlers, h...)
}

func (group *RouterGroup) addRoute(method, pattern string, handle HandlerFunc) {
	pattern = group.prefix + pattern
	group.engine.router.addRoute(method, pattern, handle)
}

func (group *RouterGroup) GET(pattern string, handle HandlerFunc) {
	group.addRoute("GET", pattern, handle)
}

func (group *RouterGroup) POST(pattern string, handle HandlerFunc) {
	group.addRoute("POST", pattern, handle)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			context.handlers = append(context.handlers, group.middlers...)
		}
	}
	engine.router.handle(context)
}
