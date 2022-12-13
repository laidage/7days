package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	handlers map[string]HandlerFunc
	root     *Node
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc), root: newNode()}
}

func parsePattern(method, pattern string) []string {
	parts := make([]string, 0)
	parts = append(parts, method)
	for _, part := range strings.Split(pattern, "/") {
		if part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}

func (router *Router) handle(c *Context) {
	parts := parsePattern(c.Method, c.Path)
	node := router.root.search(parts, 0)
	if node != nil {
		key := node.Pattern
		handle := router.handlers[key]
		handle(c)
	} else {
		c.HTML(http.StatusNotFound, "<p>404 not found.</p>")

	}
}

func (router *Router) addRoute(method, pattern string, handle HandlerFunc) {
	parts := parsePattern(method, pattern)
	key := strings.Join(parts, "-")
	router.handlers[key] = handle
	router.root.insert(parts, 0)
}
