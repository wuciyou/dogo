package controller

import (
	"github.com/wuciyou/dogo/router"
)

func init() {
	index := &indexController{}

	router.GetRouter("/indexTest", index.indexTest)

	router.GetRouter("/", index.index)
}
