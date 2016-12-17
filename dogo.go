package dogo

import (
	"net/http"
)

type dogo struct {
	serveMux *http.ServeMux
}

var DoGo *dogo

func (d *dogo) handler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	checkpipelin := Commonpipeline.each(func(pipelin *pipelineNode) bool {
		DogoLog.Debugf("Start call PipelineRun by name [%s]", pipelin.name)

		return pipelin.h.PipelineRun(response, request)
	})
	if !checkpipelin {
		return
	}
}

// start servers
func Start() {
	// 添加日志记录
	request_log := &PipelineLog{}
	Commonpipeline.AddFirst(PIPELINE_LOG, request_log)

	// 添加路由解析
	context := &pipelineContext{}
	Commonpipeline.AddLast(PIPELINE_CONTEXT, context)

	if RunTimeConfig.UserSession {
		// UserSession
		session := &pipelineSession{}
		Commonpipeline.AddFirst(PIPELINE_SESSION, session)
	}
	DogoLog.Infof("Start Dogo in the port:%s", RunTimeConfig.Port)
	DoGo.start()
}

func (t *dogo) start() {

	http.HandleFunc("/", t.handler)

	http.ListenAndServe(":"+RunTimeConfig.Port, nil)
}

func init() {
	DoGo = &dogo{serveMux: http.NewServeMux()}
}
