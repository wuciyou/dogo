package dogo

type pipelineFinishRequest struct {
}

func (f *pipelineFinishRequest) PipelineRun(ctx *Context) bool {
	ctx.AddHeader("Server", RunTimeConfig.serverName)

	data := make([]byte, ctx.response.writeBuf.Len())
	ctx.response.writeBuf.Read(data)
	ctx.response.rw.Write(data)

	DogoLog.Infof("response data:%s, byte:%v ", string(data), data)
	return true
}
