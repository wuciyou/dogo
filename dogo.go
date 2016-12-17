package dogo

import (
	"net/http"
)

type dogo struct {
	serveMux *http.ServeMux
}

var DoGo *dogo

func (d *dogo) handler(response http.ResponseWriter, request *http.Request) {

	checkpipelin := Commonpipeline.each(func(pipelin *pipelineNode) bool {
		if RunTimeConfig.IsDebug() {
			DogoLog.Printf("start call PipelineRun by name [%s]", pipelin.name)
		}

		DogoLog.Printf("pipelin.h:%+v \n", pipelin.h.PipelineRun)

		return pipelin.h.PipelineRun(response, request)
	})
	if !checkpipelin {
		return
	}

}

// start servers
func Start() {
	DoGo.start()
}

func (t *dogo) start() {

	http.HandleFunc("/", t.handler)

	http.ListenAndServe(":"+RunTimeConfig.Port, nil)
}

func init() {
	DoGo = &dogo{serveMux: http.NewServeMux()}

	// 添加日志记录
	dogo_log := &LogPipeline{}
	Commonpipeline.AddFirst(PIPELINE_LOG, dogo_log)

	// 添加路由解析
	context := &PipelineContext{}
	Commonpipeline.AddLast(PIPELINE_CONTEXT, context)

}
