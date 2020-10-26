package comment

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/display/comment/consts"
	"stream-player/src/nathan.com/video-server/display/comment/db"
)

func PostComment(c *gee.Context) {
	bodyJson, _ := ioutil.ReadAll(c.Request.Body)
	commentInfo := &consts.CommentInfo{}
	if err := json.Unmarshal(bodyJson, commentInfo); err != nil {
		c.SendErrorResponse(400, "request body parse failed")
		return
	}
	if err := db.AddNewComment(commentInfo.VideoId, commentInfo.AuthorId, commentInfo.Content); err != nil {
		c.SendErrorResponse(500, "insert into db failed")
		return
	}

	c.SendStringResponse(201, "post a comment successful")
}

func ListComments(c *gee.Context) {
	videoId, err := strconv.Atoi(c.GetQueryValue("video_id"))
	if err != nil {
		c.SendErrorResponse(400, "vid is not invalid")
	}
	from, err := strconv.Atoi(c.GetQueryValue("from"))
	if err != nil {
		c.SendErrorResponse(400, "from is not invalid")
	}
	to, err := strconv.Atoi(c.GetQueryValue("to"))
	if err != nil {
		c.SendErrorResponse(400, "to is not invalid")
	}
	commentList, err := db.ListVideoComments(videoId, from, to)
	if err != nil {
		c.SendErrorResponse(500, "read db error")
	}
	c.SendJsonResponse(200, commentList)
}
