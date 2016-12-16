package pipelines

import (
	"log"
	"net/http"
)

type LogPipeline struct {
}

func (l *LogPipeline) PipelineRun(w http.ResponseWriter, r *http.Request) bool {
	log.Printf("request:%+v \n ", r)
	return true
}
