package user

import (
	"encoding/json"
	"io/ioutil"
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/display/session"
	"stream-player/src/nathan.com/video-server/display/user/consts"
	"stream-player/src/nathan.com/video-server/display/user/db"
)

func Login(c *gee.Context) {
	username := c.GetParamValue("username")
	c.SendStringResponse(200, username)
}

func CreateUser(c *gee.Context) {
	bodyJson, _ := ioutil.ReadAll(c.Request.Body)
	userInfo := &consts.UserCredential{}
	if err := json.Unmarshal(bodyJson, userInfo); err != nil {
		c.SendErrorResponse(400, "request body parse failed")
		return
	}
	if err := db.AddUserCredential(userInfo.UserName, userInfo.PassWord); err != nil {
		c.SendErrorResponse(500, "DB ops failed")
		return
	}

	sessionId := session.CreateNewSessionId(userInfo.UserName)
	signedUp := &consts.SignedUp{Success: true, SessionId: sessionId}

	if sessionJson, err := json.Marshal(signedUp); err != nil {
		c.SendErrorResponse(500, "json dump failed")
		return
	} else {
		c.SendJsonResponse(201, sessionJson)
	}
}
