package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"stream-player/src/nathan.com/video-server/api/handlers"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", handlers.CreateUser)
	router.POST("/user/:username", handlers.Login)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}
