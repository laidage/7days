package gee

import (
	"net/http"
	"strings"
	"log"
)

type Router struct {
	handlers map[string]HandlerFunc
	root     *Node
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc), root: NewNode()}
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
	log.Printf("%v", parts)
	node := router.root.search(parts, 0)
	if node != nil {
		key := node.Pattern
		log.Printf("%v", key)
		handle := router.handlers[key]
		handle(c)
	} else {
		c.HTML(http.StatusNotFound, "<p>404 not found.</p>")

	}
}

func (router *Router) addRoute(method, pattern string, handle HandlerFunc) {
	key := method + pattern
	log.Printf("%v", key)
	router.handlers[key] = handle
	parts := parsePattern(method, pattern)
	log.Printf("%v", parts)
	router.root.insert(parts, 0)
}
