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
	body, _ := ioutil.ReadAll(c.Request.Body)
	userLoginInfo := &consts.UserCredential{}
	if err := json.Unmarshal(body, userLoginInfo); err != nil {

	}

	sessionId := session.CreateNewSessionId(username)
	signedIn := &consts.SignedUp{
		Success:   true,
		SessionId: sessionId,
	}
	c.SendJsonResponse(200, signedIn)
}

func CreateUser(c *gee.Context) {
	bodyJson, _ := ioutil.ReadAll(c.Request.Body)
	userInfo := &consts.UserCredential{}
	if err := json.Unmarshal(bodyJson, userInfo); err != nil {
		c.SendErrorResponse(400, "request body parse failed")
		return
	}
	if err := db.AddUserCredential(userInfo.UserName, userInfo.PassWord); err != nil {
		c.SendErrorResponse(500, "insert into db failed")
		return
	}

	sessionId := session.CreateNewSessionId(userInfo.UserName)
	signedUp := &consts.SignedUp{Success: true, SessionId: sessionId}

	c.SendJsonResponse(201, signedUp)
}
