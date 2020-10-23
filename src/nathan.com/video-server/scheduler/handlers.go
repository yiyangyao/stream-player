package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"stream-player/src/nathan.com/video-server/scheduler/db"
)

func videoDeleteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid, err := strconv.Atoi(p.ByName("vid"))
	if err != nil {
		sendResponse(w, 400, "vid is not invalid")
	}

	if err := db.AddVideoDeletionRecord(vid); err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}
}
