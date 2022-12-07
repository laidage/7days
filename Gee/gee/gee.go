// engine结构实现了servehttp方法，get post 添加路由，运行，servehttp，本质上是保存路由表
package gee

import (
	"fmt"
	"log"
	"net/http"
)

// router map["method-url"] = handlerfunc
type Engine struct {
	router map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]http.HandlerFunc)}
}

func (engine *Engine) addRoute(method, pattern string, handle http.HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handle
}

func (engine *Engine) GET(pattern string, handle http.HandlerFunc) {
	log.Printf("get")
	engine.addRoute("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle http.HandlerFunc) {
	log.Printf("post")
	engine.addRoute("POST", pattern, handle)
}

func (engine *Engine) Run(addr string) (err error) {
	log.Printf("run")
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("%v", req.Method)
	key := req.Method + "-" + req.URL.Path
	if engine.router[key] != nil {
		handle := engine.router[key]
		handle(w, req)
	} else {
		fmt.Fprintf(w, "404 not found.")

	}
}
