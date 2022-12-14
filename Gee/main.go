package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

// map[string] interface{}
type H map[string]interface{}

func logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Printf("logger1 start:")
		before_time := time.Now()
		c.Next()
		after_time := time.Now()
		log.Printf("logger1 use time: %v", after_time.Sub(before_time))
	}
}

func logger2() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Printf("logger2 start:")
		before_time := time.Now()
		c.Next()
		after_time := time.Now()
		log.Printf("logger2 use time: %v", after_time.Sub(before_time))
	}
}

func main() {
	engine := gee.New()
	engine.Use(logger())
	engine.GET("", helloHandle)
	engine.GET("/hello/one", helloHandle)
	engine.GET("/:lang/1", urlHandle)
	engine.GET("/hello/two", helloHandle)
	engine.GET("/nil/*", helloHandle)
	engine.GET("/:lang/2", urlHandle)
	engine.POST("/hello/three", helloHandle)
	group := engine.Group("/vvv")
	group.Use(logger2())
	group.GET("/:lang/2", urlHandle)
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
