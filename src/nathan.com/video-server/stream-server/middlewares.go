package main

import (
	"log"
	"net/http"
	"stream-player/src/nathan.com/gee-web/gee"
	"time"
)

func StreamLimiter() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		cl := NewConnLimiter(COON_LIMITER_BUFFER)

		if !cl.GetStreamConn() {
			c.SendErrorResponse(http.StatusTooManyRequests, "too many requests")
			return
		}

		// process request
		c.Next()

		cl.ReleaseStreamConn()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
