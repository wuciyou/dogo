package dogo

type PipelineLog struct {
}

func (l *PipelineLog) PipelineRun(ctx *Context) bool {

	DogoLog.Infof("request:%+v \n ", ctx.Request)
	return true
}
