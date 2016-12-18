package dogo

import (
// "net/http"
)

type pipelineRouter struct {
}

func (prouter *pipelineRouter) PipelineRun(ctx *Context) bool {
	routerContainer, _ := router.match(ctx.Request.URL.Path)
	if routerContainer == nil {
		// http.NotFound(response, request)
		return false
	}

	if string(routerContainer.method) != ctx.Request.Method {
		// http.NotFound(response, request)
		DogoLog.Warningf("Request method[%s] must be %s,url:%s", ctx.Request.Method, routerContainer.method, ctx.Request.URL.Path)
		return false
	}

	routerContainer.ch(ctx)
	return true
}
