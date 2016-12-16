package dogo

import (
	"github.com/wuciyou/dogo/pipelines"
	"net/http"
	"sync"
)

type dogo struct {
	serveMux *http.ServeMux
	mu       sync.RWMutex
	router   map[string]ContextHandle
}

var DoGo *dogo

func (d *dogo) setRouter(pattern string, ch ContextHandle) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	d.router[pattern] = ch
}

func (d *dogo) getRouter(pattern string) ContextHandle {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if c, ok := d.router[pattern]; ok {
		return c
	} else {
		return nil
	}
}

func (d *dogo) Router(pattern string, ch ContextHandle) {

	if d.getRouter(pattern) != nil {
		DogoLog.Panicf("The pattern[%s] already exists in the Router", pattern)
	}
	d.setRouter(pattern, ch)

}

func (d *dogo) handler(response http.ResponseWriter, request *http.Request) {
	ch := d.getRouter(request.URL.Path)

	if ch == nil {
		http.NotFound(response, request)
		return
	}
	if RunTimeConfig.IsDebug() {
		DogoLog.Printf("pattern:%s \n ", request.URL.Path)
	}

	checkpipelin := Commonpipeline.each(func(pipelin *pipelineNode) bool {
		if RunTimeConfig.IsDebug() {
			DogoLog.Printf("start call PipelineRun by name [%s]", pipelin.name)
		}
		return pipelin.h.PipelineRun(response, request)
	})
	if !checkpipelin {
		return
	}
	context := &Context{}
	context.parse(response, request)
	ch(context)

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
	if RunTimeConfig.IsDebug() {
		DogoLog.Printf("ThinkGo default config: %+v \n ", RunTimeConfig)
	}
}

func init() {
	DoGo = &dogo{serveMux: http.NewServeMux(), router: make(map[string]ContextHandle)}
}
