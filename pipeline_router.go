package dogo

import (
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
)

type pipelineRouter struct {
}

func (prouter *pipelineRouter) PipelineRun(ctx *context.Context) bool {
	routerContainer, pattern := router.match(ctx.Request.URL.Path)
	ctx.Pattern = pattern
	if routerContainer == nil {
		ctx.NotFound()
		return false
	}

	if routerContainer.method != "" && string(routerContainer.method) != ctx.Request.Method {
		// http.NotFound(response, request)
		dglog.Warningf("Request method[%s] must be %s,url:%s", ctx.Request.Method, routerContainer.method, ctx.Request.URL.Path)
		return false
	}

	routerContainer.ch(ctx)
	return true
}
