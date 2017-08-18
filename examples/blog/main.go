package main

import (
	"github.com/wuciyou/dogo"
)

func main() {
	dogo.Route.Get("/hello/wuciyou", func(c *dogo.Context) {
		dogo.Debug("有新的页面请求")
		c.W.Write([]byte("hello get"))
	})
	dogo.Route.Post("/hello/wuciyou", func(c *dogo.Context) {
		dogo.Debug("有新的页面请求")
		c.W.Write([]byte("hello post"))
	})
	dogo.App().Run()
	dogo.App().Run()
}
