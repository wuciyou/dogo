package controller

import (
	"github.com/wuciyou/dogo/router"
)

func init() {
	redirect := &redirectController{}

	router.GetRouter("/", redirect.do)
}
