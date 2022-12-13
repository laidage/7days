package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	W          http.ResponseWriter
	Req        *http.Request
	Method     string
	StatusCode int
	Path       string
	index      int
	handlers   []HandlerFunc
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		W:        w,
		Req:      req,
		Path:     req.URL.Path,
		Method:   req.Method,
		index:    -1,
		handlers: make([]HandlerFunc, 0),
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.PostFormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.W.Header().Set(key, value)
}

func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)
	encoder.Encode(obj)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	fmt.Fprintf(c.W, html)
}

func (c *Context) String(code int, format string, values ...string) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	fmt.Fprintf(c.W, format, values)
}
