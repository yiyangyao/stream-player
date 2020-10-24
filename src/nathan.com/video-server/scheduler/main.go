package main

import (
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/scheduler/worker"
)

func main() {
	//c := make(chan int)
	go worker.Start()

	app := gee.Default()
	clearGroup := app.Group("/clear")
	clearGroup.POST("/video-delete-record/:vid", videoDeleteHandler)

	//<- c
	_ = app.Run(":9001")
}
