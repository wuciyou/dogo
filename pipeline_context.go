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

type PipelineContext struct {
}

func (c *PipelineContext) PipelineRun(response http.ResponseWriter, request *http.Request) bool {
	ch, _ := router.match(request.URL.Path)
	if ch == nil {
		http.NotFound(response, request)
		return false
	}

	context := &Context{}
	context.parse(response, request)
	ch(context)
	return true
}
