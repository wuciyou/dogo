package handle

import (
	"github.com/wuciyou/dogo/config"

	"github.com/wuciyou/dogo/context"

	"github.com/wuciyou/dogo/dglog"
)

type FinishRequest struct {
}

func (f *FinishRequest) PipelineRun(ctx *context.Context) bool {

	serverName, err := config.GetString("SERVER_NAME")
	if err != nil {
		dglog.Error(err)
	}
	ctx.AddHeader("Server", serverName)
	data := ctx.Flush(true)
	dglog.Infof("response data:%s, byte:%v ", string(data), data)
	return true
}
