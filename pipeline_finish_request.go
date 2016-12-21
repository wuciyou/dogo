package dogo

import (
	"github.com/wuciyou/dogo/config"
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
)

type pipelineFinishRequest struct {
}

func (f *pipelineFinishRequest) PipelineRun(ctx *context.Context) bool {
	ctx.AddHeader("Server", config.RunTimeConfig.ServerName())
	data := ctx.Flush(true)
	dglog.Infof("response data:%s, byte:%v ", string(data), data)
	return true
}
