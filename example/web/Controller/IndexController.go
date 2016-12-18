package Controller

import (
	"github.com/wuciyou/dogo"
)

type IndexController struct {
	Name string
	Age  int
}

func (index *IndexController) Index(c *dogo.Context) {

	name := make(map[string]string)
	name["name"] = "wuciyou"
	c.SetHeader("Server", "wuciyou")

	c.Assign(name)
	c.Display("View/index.html")

}

func init() {
	index := &IndexController{}
	dogo.GetRouter("/", index.Index)
}
