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
	routerContainer, _ := router.match(request.URL.Path)
	if routerContainer == nil {
		http.NotFound(response, request)
		return false
	}

	if string(routerContainer.method) != request.Method {
		http.NotFound(response, request)
		DogoLog.Warningf("Request method[%s] must be %s,url:%s", request.Method, routerContainer.method, request.URL.Path)
		return false
	}

	context := &Context{}
	context.parse(response, request)
	routerContainer.ch(context)
	return true
}
