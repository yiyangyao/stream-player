package main

import (
	"stream-player/src/nathan.com/gee-web/gee"
	"stream-player/src/nathan.com/video-server/display/middlerwares"
	"stream-player/src/nathan.com/video-server/display/user"
)

func main() {
	app := gee.Default()
	app.Apply(middlerwares.ValidateSession())

	userGroup := app.Group("/user")
	{
		userGroup.POST("/add", user.CreateUser)
		userGroup.POST("/login/:username", user.Login)
	}

	_ = app.Run(":9000")
}
