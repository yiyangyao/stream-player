package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"stream-player/src/nathan.com/gee-web/gee"
)

func streamHandler(c *gee.Context) {
	fileName := c.GetParamValue("video-name")
	filePath := VIDEO_DIR + fileName

	video, err := os.Open(filePath)
	if err != nil {
		log.Printf("open video file failed: %v", err)
		c.SendErrorResponse(http.StatusInternalServerError, "open video file failed， please check the video file")
		return
	}

	c.SendVideoMP4Response(video)

	defer video.Close()
}

func uploadHandler(c *gee.Context) {
	c.Request.Body = http.MaxBytesReader(c.Response, c.Request.Body, MAX_UPLOAD_SIZE)
	if err := c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		c.SendErrorResponse(http.StatusBadRequest, "video size exceeds limit")
		return
	}
	file, _, err := c.Request.FormFile("file") // <form name="file">
	if err != nil {
		c.SendErrorResponse(http.StatusInternalServerError, "internal server error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("read file error %v", err)
		c.SendErrorResponse(http.StatusInternalServerError, "internal server error: read file error")
		return
	}

	fileName := c.GetParamValue("video-name")
	if err := ioutil.WriteFile(VIDEO_DIR+fileName, data, 0666); err != nil {
		log.Printf("write file content error: %v", err)
		c.SendErrorResponse(http.StatusInternalServerError, "internal server error: write file content error")
		return
	}

	c.SendStringResponse(http.StatusCreated, "upload successfully")
}

func uploadPageHandler(c *gee.Context) {
	t, _ := template.ParseFiles(TEMPLATE_DIR + "upload.html")
	t.Execute(c.Response, nil)
}
