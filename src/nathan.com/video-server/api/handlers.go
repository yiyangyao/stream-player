package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"stream-player/src/nathan.com/video-server/api/db"
	"stream-player/src/nathan.com/video-server/api/defs"
	"stream-player/src/nathan.com/video-server/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	body := &defs.UserCredential{}
	if err := json.Unmarshal(res, body); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := db.AddUserCredential(body.UserName, body.PassWord); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	sessionId := session.CreateNewSessionId(body.UserName)
	success := &defs.SignedUp{Success: true, SessionId: sessionId}

	if res, err := json.Marshal(success); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
		return
	} else {
		sendNormalResponse(w, string(res), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	io.WriteString(w, username)
}
