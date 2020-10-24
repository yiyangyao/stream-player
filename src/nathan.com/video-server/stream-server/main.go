package main

import (
	"stream-player/src/nathan.com/gee-web/gee"
)

func main() {
	app := gee.Default()

	// create group v2
	streamGroup := app.Group("/stream")
	// apply middleware for group v2
	streamGroup.Apply(StreamLimiter())
	{
		streamGroup.GET("/videos/:video-name", streamHandler)
		streamGroup.POST("/upload/:video-name", uploadHandler)
		streamGroup.GET("/upload-page", uploadPageHandler)
	}

	app.Run(":9999")
}
