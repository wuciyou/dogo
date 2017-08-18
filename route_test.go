package dogo

import (
	"strings"
	"testing"
)

func TestAddRoute(t *testing.T) {

	Route.Get("/hello/wuciyou/aa/bb/cc", func(c *Context) {})
	Route.Get("/hello/wuciyou1", func(c *Context) {})
	Route.Get("/hello/wuciyou2", func(c *Context) {})
	Route.Post("/hello/wuciyou3", func(c *Context) {})

	h := Route.checkMethod("POST", strings.Split("/hello/wuciyou3", "/"))
	t.Log(h)

	t.Logf("Route %+v", Route)
	t.Log("hello word")
}
