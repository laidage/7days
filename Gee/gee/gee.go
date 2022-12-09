// engine结构实现了servehttp方法，get post 添加路由，运行，servehttp，本质上是保存路由表
package gee

import (
	"log"
	"net/http"
)

// router map["method-url"] = handlerfunc
type Engine struct {
	router *Router
}

type HandlerFunc func(c *Context)

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method, pattern string, handle HandlerFunc) {
	key := method + "-" + pattern
	engine.router.handlers[key] = handle
}

func (engine *Engine) GET(pattern string, handle HandlerFunc) {
	log.Printf("get")
	engine.addRoute("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandlerFunc) {
	log.Printf("post")
	engine.addRoute("POST", pattern, handle)
}

func (engine *Engine) Run(addr string) (err error) {
	log.Printf("run")
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	engine.router.handle(context)
}
