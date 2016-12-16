package dogo

import (
	"github.com/wuciyou/dogo/pipelines"
	"log"
	"net/http"
)

type dogo struct {
	serveMux *http.ServeMux
}

var DoGo *dogo

func (t *dogo) Router(pattern string, controller dogoController) {
	t.serveMux.HandleFunc(pattern, controller.handler)
}

func (t *dogo) handler(response http.ResponseWriter, request *http.Request) {
	_, pattern := t.serveMux.Handler(request)

	if RunTimeConfig.IsDebug() {
		log.Printf("pattern:%s \n ", pattern)
	}
	log.Printf("Commonpipeline:%+v \n", Commonpipeline)
	Commonpipeline.each(func(pipelin *pipelineNode) bool {
		log.Printf("start call PipelineRun by name [%s] \n ", pipelin.name)
		return pipelin.h.PipelineRun(response, request)
	})

	t.serveMux.ServeHTTP(response, request)
}

// start servers
func (t *dogo) Start() {
	t.start()
}

func (t *dogo) start() {
	initConfig()

	http.HandleFunc("/", t.handler)

	dogo_log := &pipelines.LogPipeline{}
	Commonpipeline.AddFirst("dogo_log", dogo_log)

	http.ListenAndServe(":"+RunTimeConfig.Port, nil)
}

func initConfig() {
	if RunTimeConfig.Port == "" {
		RunTimeConfig.Port = "8080"
	}
	if RunTimeConfig.IsDebug() {
		log.Printf("ThinkGo default config: %+v \n ", RunTimeConfig)
	}
}

func init() {
	DoGo = &dogo{serveMux: http.NewServeMux()}
}
