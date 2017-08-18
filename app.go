package dogo

import (
	"fmt"
	"net/http"
)

type appServe struct {
	serveMux *http.ServeMux
	conf     dogoConfig
}

func App() *appServe {
	var app = &appServe{conf: defaultConfig, serveMux: &http.ServeMux{}}
	return app
}

func (app *appServe) Run() {
	var addr = fmt.Sprintf("%s:%d", app.conf.String("SERVER", "WEB_IP"), app.conf.Int("SERVER", "WEB_PORT"))

	Dglog.Infof("Starting .... %s", addr)

	// Register root route handle func
	app.serveMux.HandleFunc("/", app.do)

	// Listening on addr
	http.ListenAndServe(addr, app.serveMux)
}

func (app *appServe) do(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Dglog.Debugf("pattern:%s", r.Form.Get("name"))

	h := Route.checkRoute(r)
	if h == nil {
		Dglog.Errorf("Not found page :%s", r.RequestURI)
	} else {
		h(&Context{R: r, W: w})
	}
}
