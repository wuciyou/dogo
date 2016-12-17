package dogo

import (
	"net/http"
)

type LogPipeline struct {
}

func (l *LogPipeline) PipelineRun(w http.ResponseWriter, r *http.Request) bool {
	DogoLog.Infof("request:%+v \n ", r)
	return true
}
