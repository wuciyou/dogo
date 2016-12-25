package handle

import (
	"github.com/wuciyou/dogo/common"
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
	"github.com/wuciyou/dogo/hooks"
	"github.com/wuciyou/dogo/router"
)

type Router struct {
}

func (prouter *Router) PipelineRun(ctx *context.Context) bool {
	hooks.Listen(common.ROUTER_PARSE_BEGIN, ctx)
	contextHandle, pattern, err := router.Match(ctx.Request)
	ctx.Pattern = pattern
	if err != nil {
		ctx.NotFound()
		dglog.Error(err)
		return false
	}
	if contextHandle == nil {
		ctx.NotFound()
		return false
	}

	contextHandle(ctx)
	return true
}
