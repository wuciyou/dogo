package thinkgo

import (
	"net/http"
)

type thinkController interface {
	handler(http.ResponseWriter, *http.Request)
}

type BaseController struct {
	Get  map[string][]string
	Post map[string][]string
	// Get and Post merge
	Request map[string][]string
}

func (c *BaseController) handler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("wuciyou hello word"))
}
