package session

import (
	"fmt"
	"github.com/wuciyou/dogo/session/handle"
	"testing"
)

var sessionContainer *SessionContainer

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

func TestSession(t *testing.T) {
	inits(t)
	sid, _ := GenerateSid()
	sid = "9e76c6af-99fe-8adf-db47-cd288cb809b9"
	sessionContainer = GetSession(sid)
	for i := 0; i <= 0; i++ {
		sessionStr(t, i)
		sessionMap(t, i)
	}
}

func sessionStr(t *testing.T, i int) {
	addValue := fmt.Sprintf("index:%d%d", i, i)
	name := fmt.Sprintf("name_str:%d", i)
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

func sessionMap(t *testing.T, i int) {
	addValue := make(map[string]string)
	addValue["key1"] = "map1"
	addValue["key2"] = "map2"
	name := fmt.Sprintf("name_map:%d", i)

	sessionContainer.Add(name, addValue)
	getMapValue := make(map[string]string)
	sessionContainer.Get(name, &getMapValue)
	t.Logf("getMapValue:%+v", getMapValue)
	for k, v := range addValue {
		if v != getMapValue[k] {
			t.Errorf("Session get name:%s error: addValue:%s neq getValue:%s ", name, v, getMapValue[k])
		}
	}

	sessionContainer.Delete(name)
	getMapValue = make(map[string]string)
	sessionContainer.Get(name, &getMapValue)
	if len(getMapValue) > 0 {
		t.Errorf("Session name:%s is not null", name)
	}
}
