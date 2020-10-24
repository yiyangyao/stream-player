package middlerwares

import (
	"log"
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/display/user"
	"time"
)

func ValidateSession() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()

		user.ValidateUserSession(c.Request)
		// process request
		c.Next()

		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
