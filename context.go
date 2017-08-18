package dogo

import (
	"net/http"
)

type Context struct {
	R *http.Request
	W http.ResponseWriter
}
