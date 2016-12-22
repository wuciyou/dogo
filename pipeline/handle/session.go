package handle

import (
	"github.com/wuciyou/dogo/context"
)

type Session struct {
}

func (s *Session) PipelineRun(ctx *context.Context) bool {

	// cookie, err := request.Cookie(RunTimeConfig.SessionName)
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		DogoLog.Debug("SessionName not present in the cookie ")
	// 		SessionNameid := "wuciyourqyqqqq"
	// 		DogoLog.Debugf("Reset SessionName[%s] to cookie ", SessionNameid)

	// 		newCookie := &http.Cookie{}
	// 		newCookie.Name = RunTimeConfig.SessionName
	// 		newCookie.Value = SessionNameid

	// 		http.SetCookie(response, newCookie)

	// 	}
	// }

	// DogoLog.Debugf("Read SessionName success cookie:%+v", cookie)

	return true
}
