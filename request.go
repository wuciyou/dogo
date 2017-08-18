package dogo

import (
	"net/http"
)

type Request struct {
	*http.Request
}
