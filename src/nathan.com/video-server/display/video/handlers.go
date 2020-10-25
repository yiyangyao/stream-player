package video

import (
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/display/user"
	userDB "stream-player/src/nathan.com/video-server/display/user/db"
	videoDB "stream-player/src/nathan.com/video-server/display/video/db"
	schedulerDB "stream-player/src/nathan.com/video-server/scheduler/db"
)

func CreateVideo(c *gee.Context) {
	username := c.Request.Header.Get(user.HEADER_FIELD_UNAME)
	userInfo, err := userDB.GetUserCredential(username)
	if userInfo == nil || err != nil {
		c.SendErrorResponse(500, "user is not existed")
	}
	video, err := videoDB.AddNewVideo(userInfo.UserID, c.GetParamValue("video-name"))
	if err != nil {
		c.SendErrorResponse(500, "insert db failed")
	}
	c.SendJsonResponse(201, video)
}

func GetVideo(c *gee.Context) {
	videoName := c.GetQueryValue("video-name")
	videoInfo, err := videoDB.GetVideoInfo(videoName)
	if videoInfo == nil || err != nil {
		c.SendErrorResponse(500, "video is not existed")
	}
	c.SendJsonResponse(201, videoInfo)
}

func DeleteVideo(c *gee.Context) {
	videoName := c.GetParamValue("video-name")
	videoInfo, err := videoDB.GetVideoInfo(videoName)
	if videoInfo == nil || err != nil {
		c.SendErrorResponse(500, "video is not existed")
	}
	if err := videoDB.DeleteVideo; err != nil {
		c.SendErrorResponse(500, "delete db failed")
	}
	if err := schedulerDB.AddVideoDeletionRecord(videoInfo.VideoId); err != nil {
		c.SendErrorResponse(500, "delete db failed")
	}
	c.SendJsonResponse(201, videoInfo)
}

func ListVideos(c *gee.Context) {
	username := c.Request.Header.Get(user.HEADER_FIELD_UNAME)
	userInfo, err := userDB.GetUserCredential(username)
	if userInfo == nil || err != nil {
		c.SendErrorResponse(500, "user is not existed")
	}
	videoList, err := videoDB.ListVideoInfo(userInfo.UserID)
	if err != nil {
		c.SendErrorResponse(500, "get video list failed")
	}
	c.SendJsonResponse(200, videoList)
}
