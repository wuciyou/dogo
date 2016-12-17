package Controller

import (
	"github.com/wuciyou/dogo"
)

type IndexController struct {
}

func (index *IndexController) Index(c *dogo.Context) {

	var body = " <center> Hello Word! <br> This is a Dogo page  <center> "

	c.Response.Write([]byte(body))
}

func init() {
	index := &IndexController{}
	dogo.GetRouter("/", index.Index)
}
