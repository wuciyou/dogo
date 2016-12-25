package handle

// import (
// 	"github.com/wuciyou/dogo/config"
// 	"github.com/wuciyou/dogo/context"
// 	"github.com/wuciyou/dogo/dglog"
// 	"github.com/wuciyou/dogo/session"
// 	"net/http"
// )

// type Session struct {
// }

// func (s *Session) PipelineRun(ctx *context.Context) bool {
// 	sessionName, err := config.GetString("SESSION.NAME")
// 	if err != nil {
// 		dglog.Error(err)
// 	}

// 	cookie, err := request.Cookie(sessionName)
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			sid, err := session.GenerateSid()
// 			if err != nil {
// 				dglog.Error(err)
// 			}

// 			DogoLog.Debugf("Reset SessionName[%s] to cookie ", sid)

// 			newCookie := &http.Cookie{}
// 			newCookie.Name = sessionName
// 			newCookie.Value = sid

// 			http.SetCookie(response, newCookie)

// 		}
// 	}
// 	return true
// }
