package session

import (
	"fmt"
	"github.com/wuciyou/dogo/session/handle"
	"testing"
)

func inits(t *testing.T) {
	fs := handle.FileStoreEntity
	err := fs.Open()
	if err != nil {

		t.Errorf("Could not call FileStoreEntity open menthod:%v ", err)
	}
	SetSessionStore(fs)
}

func TestAdd(t *testing.T) {

	// Add("wuciyou")

}

func TestGenerateSid(t *testing.T) {
	for i := 0; i <= 20; i++ {
		_, err := GenerateSid()
		if err != nil {
			t.Errorf("Could not GenerateSid %v", err)
		}
	}
}

func TestGetSession(t *testing.T) {
	inits(t)
	sid, _ := GenerateSid()
	sessionContainer := GetSession(sid)
	for i := 0; i <= 10; i++ {
		addValue := fmt.Sprintf("index:%d%d", i, i)
		name := fmt.Sprintf("name_wuciyou:%d", i)
		sessionContainer.Add(name, addValue)
		var getValue string
		sessionContainer.Get(name, &getValue)
		if getValue != addValue {
			t.Errorf("Session get name:%s error: addValue:%s neq getValue:%s ", name, addValue, getValue)
		}
		sessionContainer.Delete(name)
		var getValue1 string
		sessionContainer.Get(name, &getValue1)
		if getValue1 == addValue {
			t.Errorf("Session get name:%s error: addValue:%s eq getValue1:%s ", name, addValue, getValue1)
		}

	}

}
