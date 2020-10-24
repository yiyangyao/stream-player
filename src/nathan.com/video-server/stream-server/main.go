package main

import (
	"stream-player/src/nathan.com/gee-web/gee"
)

func main() {
	// create gee engine with logger and recover
	app := gee.Default()
	// create group stream
	streamGroup := app.Group("/stream")
	// apply StreamLimiter middleware for group stream
	streamGroup.Apply(StreamLimiter())
	{
		streamGroup.GET("/videos/:video-name", streamHandler)
		streamGroup.POST("/upload/:video-name", uploadHandler)
		streamGroup.GET("/upload-page", uploadPageHandler)
	}

	_ = app.Run(":9999")
}
