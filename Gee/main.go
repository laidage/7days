package main

import (
	"gee"
	"net/http"
)

// map[string] interface{}
type H map[string]interface{}

func main() {
	engine := gee.New()
	engine.GET("", helloHandle)
	engine.GET("/hello/one", helloHandle)
	engine.GET("/:lang/1", urlHandle)
	engine.GET("/hello/two", helloHandle)
	engine.GET("/:lang/2", urlHandle)
	engine.POST("/hello/three", helloHandle)
	engine.Run(":9999")
}

func helloHandle(c *gee.Context) {
	c.HTML(http.StatusOK, "Hello, nihao")
}

func urlHandle(c *gee.Context) {
	// fmt.Fprintf(w, "this is you url: %v", req.URL.Path)
	url := H{
		"url": c.Path,
	}
	c.Json(http.StatusOK, url)
}
