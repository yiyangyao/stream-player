package main

import (
	"fmt"
	"net/http"
	"stream-player/src/nathan.com/gee-web/gee"
)

func main() {
	//app := gee.New()
	app := gee.Default()
	// apply global middleware
	//app.Apply(gee.Logger())

	// add global router
	app.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	// panic
	app.GET("/panic", func(c *gee.Context) {
		names := []string{"nathan"}
		c.String(http.StatusOK, names[100])
	})

	// create group v1
	v1 := app.Group("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=nathan
			data := fmt.Sprintf("hello %s, you're at %s\n", c.GetQueryValue("name"), c.Path)
			c.String(http.StatusOK, data)
		})

		v1.POST("/login", func(c *gee.Context) {
			c.Json(200, gee.H{
				"username": c.GetPostFormValue("username"),
				"password": c.GetPostFormValue("password"),
			})
		})
	}

	// create group v2
	v2 := app.Group("/v2")
	// apply middleware for group v2
	v2.Apply(gee.ForGroupV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// except /hello/nathan
			data := fmt.Sprintf("hello %s, you're at %s\n", c.GetParamValue("name"), c.Path)
			c.String(http.StatusOK, data)
		})

		v2.GET("/assets/*filepath", func(c *gee.Context) {
			c.Json(http.StatusOK, gee.H{"filepath": c.GetParamValue("filepath")})
		})
	}

	app.Run(":9999")
}

/*
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

$ curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
{"password":"1234","username":"geektutu"}

$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx

$ curl "http://localhost:9999/hello/geektutu"
hello geektutu, you're at /hello/geektutu

$ curl "http://localhost:9999/assets/css/geektutu.css"
{"filepath":"css/geektutu.css"}

$ curl "http://localhost:9999/v1/hello?name=geektutu"
hello geektutu, you're at /v1/hello

$ curl "http://localhost:9999/v2/hello/geektutu"
hello geektutu, you're at /hello/geektutu

$ curl "http://localhost:9999/"
>>> log
2019/08/17 01:37:38 [200] / in 3.14µs

(2) global + group middleware
$ curl http://localhost:9999/v2/hello/geektutu
>>> log
2019/08/17 01:38:48 [200] /v2/hello/geektutu in 61.467µs for group v2
2019/08/17 01:38:48 [200] /v2/hello/geektutu in 281µs
*/
