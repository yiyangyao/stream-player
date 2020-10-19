package main

import (
	"net/http"
	"stream-player/src/nathan.com/video-server/api/defs"
	"stream-player/src/nathan.com/video-server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-Session-Name"

func vaildateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	username, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, username)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.Header.Get(HEADER_FIELD_UNAME)
	if len(username) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
