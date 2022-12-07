package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/hello", helloHandle)
	engine.POST("/url", urlHandle)
	engine.Run(":9999")
}

func helloHandle(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, nihao")
}

func urlHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "this is you url: %v", req.URL.Path)
}
