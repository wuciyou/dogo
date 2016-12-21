package dogo

import (
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
)

type PipelineLog struct {
}

func (l *PipelineLog) PipelineRun(ctx *context.Context) bool {

	dglog.Infof("request:%+v \n ", ctx.Request)
	return true
}
