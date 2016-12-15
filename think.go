package thinkgo

import (
	"log"
	"net/http"
)

type think struct {
	serveMux *http.ServeMux
}

var ThinkGo *think

func (t *think) Router(pattern string, controller thinkController) {
	t.serveMux.HandleFunc(pattern, controller.handler)
}

func (t *think) handler(response http.ResponseWriter, request *http.Request) {
	_, pattern := t.serveMux.Handler(request)

	if RunTimeConfig.IsDebug() {
		log.Printf("pattern:%s \n ", pattern)
	}

	t.serveMux.ServeHTTP(response, request)
}

// start servers
func (t *think) Start() {
	t.start()
}

func (t *think) start() {
	initConfig()

	http.HandleFunc("/", t.handler)
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
	ThinkGo = &think{serveMux: http.NewServeMux()}
}
