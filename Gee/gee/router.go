package gee

import (
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
	root     *Node
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc), root: NewNode()}
}

func (router *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if router.handlers[key] != nil {
		handle := router.handlers[key]
		handle(c)
	} else {
		c.HTML(http.StatusNotFound, "<p>404 not found.</p>")

	}
}

func (router *Router) addRoute(method, pattern string, handle HandlerFunc) {
	key := method + "-" + pattern
	router.handlers[key] = handle
}
