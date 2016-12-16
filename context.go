package dogo

import (
	"net/http"
)

type ContextHandle func(c *Context)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (c *Context) parse(response http.ResponseWriter, request *http.Request) {
	c.Response = response
	c.Request = request
}
