package gee

import (
	"log"
	"net/http"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recover!")
				c.HTML(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
