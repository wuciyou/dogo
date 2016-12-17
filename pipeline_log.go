package dogo

import (
	"net/http"
)

type PipelineLog struct {
}

func (l *PipelineLog) PipelineRun(w http.ResponseWriter, r *http.Request) bool {
	DogoLog.Infof("request:%+v \n ", r)
	return true
}
