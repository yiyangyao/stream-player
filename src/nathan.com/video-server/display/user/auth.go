package user

import (
	"net/http"
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/display/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-Session-Name"

func ValidateUserSession(r *http.Request) bool {
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

func ValidateUser(c *gee.Context) bool {
	username := c.Request.Header.Get(HEADER_FIELD_UNAME)
	if len(username) == 0 {
		c.SendErrorResponse(403, "user auth failed")
		return false
	}
	return true
}
