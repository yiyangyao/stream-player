package main

import (
	"strconv"
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/scheduler/db"
)

func videoDeleteHandler(c *gee.Context) {
	vid, err := strconv.Atoi(c.GetParamValue("vid"))
	if err != nil {
		c.SendErrorResponse(400, "vid is not invalid")
	}

	if err := db.AddVideoDeletionRecord(vid); err != nil {
		c.SendErrorResponse(500, "Internal server error")
		return
	}
}
