package gee

import (
	"fmt"
	"log"
	"time"
)

// global middleware
func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		// process request
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
		fmt.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

// apply group v2 middleware
func ForGroupV2() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		fmt.Printf("call api %s for group v2, call time %v", c.Request.RequestURI, time.Since(t))
		c.Fail(500, "Internal Server SendErrorResponse")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
		fmt.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
