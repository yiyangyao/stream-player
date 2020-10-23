package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"stream-player/src/nathan.com/video-server/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/video-delete-record/:vid", videoDeleteHandler)

	return router
}

func main() {
	//c := make(chan int)
	go taskrunner.Start()
	r := RegisterHandlers()
	//<- c
	_ = http.ListenAndServe(":9001", r)
}
