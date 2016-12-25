package context

import (
	"github.com/wuciyou/dogo/config"
	"github.com/wuciyou/dogo/dglog"
	"github.com/wuciyou/dogo/session"
	"net/http"
)

func (c *Context) GetSession() *session.SessionContainer {
	if c.sessionContainer == nil {
		c.initSession()
	}
	return c.sessionContainer
}

func (c *Context) initSession() {
	sessionName, err := config.GetString("SESSION.NAME")
	if err != nil {
		dglog.Error(err)
	}
	cookie, err := c.Request.Cookie(sessionName)
	if err != nil {
		if err == http.ErrNoCookie {
			sid, err := session.GenerateSid()
			if err != nil {
				dglog.Error(err)
			}

			dglog.Debugf("Reset SessionName[%s] to cookie ", sid)

			c.sessionContainer = session.GetSession(sid)

			newCookie := &http.Cookie{}
			newCookie.Name = sessionName
			newCookie.Value = sid

			http.SetCookie(c.response.rw, newCookie)
		}
	} else {
		c.sessionContainer = session.GetSession(cookie.Value)
	}

}
