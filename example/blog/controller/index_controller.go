package controller

import (
	"github.com/wuciyou/dogo/context"
	"time"
)

type indexController struct {
}

func (this *indexController) index(c *context.Context) {
	c.Display("index_index.html")
}

func (this *indexController) indexTest(c *context.Context) {

	time.Sleep(time.Minute)

	c.Display("index_index.html")
}
