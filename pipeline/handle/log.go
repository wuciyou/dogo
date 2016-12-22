package handle

import (
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
)

type Log struct {
}

func (l *Log) PipelineRun(ctx *context.Context) bool {

	dglog.Infof("request:%+v \n ", ctx.Request)
	return true
}
