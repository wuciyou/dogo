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
	c.AddHeader("Server", "wuciyou")

	c.Assign(name)
	c.Display("View/index.html")

}

func (index *IndexController) ReturnJson(c *dogo.Context) {
	index.Name = "ReturnJson"
	index.Age = 25
	c.AjaxReturn(index)
}

func (index *IndexController) ReturnXml(c *dogo.Context) {
	index.Name = "ReturnXml"
	index.Age = 25
	c.AjaxReturn(index, "XML")
}

func init() {
	index := &IndexController{}
	dogo.GetRouter("/", index.Index)
	dogo.GetRouter("/returnJson", index.ReturnJson)
	dogo.GetRouter("/returnXml", index.ReturnXml)
}
