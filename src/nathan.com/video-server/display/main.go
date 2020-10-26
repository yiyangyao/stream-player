package main

import (
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/display/comment"
	"stream-player/src/nathan.com/video-server/display/middlerwares"
	"stream-player/src/nathan.com/video-server/display/user"
	"stream-player/src/nathan.com/video-server/display/video"
)

func main() {
	app := gee.Default()

	userGroup := app.Group("/user")
	{
		userGroup.POST("/register", user.CreateUser)
		userGroup.POST("/login/:username", user.Login)
	}

	videoGroup := app.Group("/video")
	videoGroup.Apply(middlerwares.ValidateSession())
	{
		videoGroup.POST("/add/:video-name", video.CreateVideo)
		videoGroup.POST("/delete/:video-name", video.DeleteVideo)
		videoGroup.GET("/detail", video.GetVideo)
		videoGroup.GET("/list", video.ListVideos)
	}

	commentGroup := app.Group("/comment")
	{
		commentGroup.GET("/list", comment.ListComments)
		commentGroup.POST("/add", comment.PostComment)
	}

	_ = app.Run(":9000")
}
