package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VideoDir + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("err %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "open video failed")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "max size")
		return
	}
	file, _, err := r.FormFile("file") // <form name="file">
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "internal err")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("read file error %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal err")
		return
	}

	fn := p.ByName("vid-id")
	if err := ioutil.WriteFile(VideoDir+fn, data, 0666); err != nil {
		log.Printf("write file err %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal err")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload successfully")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("/Users/bytedance/stream-player/videos/upload.html")
	t.Execute(w, nil)
}
