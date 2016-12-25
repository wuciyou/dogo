package Controller

import (
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/router"
)

type sessionController struct {
	Name string
	Age  string
}

func (s *sessionController) Get(c *context.Context) {
	sessionName := c.Request.FormValue("name")
	session := c.GetSession()

	var getSeesion sessionController

	session.Get(sessionName, &getSeesion)
	c.AjaxReturn(getSeesion)
}

func (s *sessionController) Add(c *context.Context) {
	sessionName := c.Request.FormValue("name")
	session := c.GetSession()

	var getSeesion = &sessionController{}

	getSeesion.Name = sessionName
	getSeesion.Age = c.Request.Form.Get("age")
	session.Add(sessionName, getSeesion)
	c.WriteString("ok")
}

func init() {
	sessionController := &sessionController{}
	router.GetRouter("/session/get", sessionController.Get)
	router.GetRouter("/session/add", sessionController.Add)
}
