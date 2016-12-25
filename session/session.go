package session

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wuciyou/dogo/dglog"
	"github.com/wuciyou/dogo/session/handle"
	"os"
	"sync"
	"time"
)

type SessionHandle interface {
	Open() error
	Close()
	Read(sid string) []byte
	Write(sid string, data []byte)
	Delete(sid string)
	Gc()
}

type SessionContainer struct {
	m   *sync.Mutex
	sid string
	// 过期时间
	expire int64
}

type sessioneManager struct {
	handle    SessionHandle
	m         *sync.Mutex
	container map[string]*SessionContainer
}

var manager *sessioneManager

func (container *SessionContainer) Add(name string, data interface{}) {
	manager.add(container.sid, name, data)
}

func (container *SessionContainer) Get(name string, data interface{}) {
	manager.get(container.sid, name, data)
}

func (container *SessionContainer) Delete(name string) {
	manager.delete(container.sid, name)
}

func (manager *sessioneManager) add(sid string, name string, data interface{}) {
	var dataBf bytes.Buffer
	enc := gob.NewEncoder(&dataBf)

	h := manager.getHandle()
	dataMap := make(map[string][]byte)
	oldData := h.Read(sid)
	if oldData != nil {
		json.Unmarshal(oldData, &dataMap)
		dglog.Debugf("session [%s] parse:%s", sid, dataMap)
	}

	if err := enc.Encode(data); err != nil {
		dglog.Errorf("Session encode fail. sid:%s, name:%s, data:%v", sid, name, data)
	}
	dataMap[name] = dataBf.Bytes()
	dglog.Debugf("Session encoder:sid:%s, name:%s, data:%v \n", sid, name, dataMap[name])
	newData, err := json.Marshal(dataMap)

	if err != nil {
		dglog.Errorf("Can't marshal session[sid:%s, name:%s, data:%v] to session ", sid, name, data)
		return
	}
	h.Write(sid, newData)
}

func (manager *sessioneManager) get(sid string, name string, data interface{}) {
	h := manager.getHandle()
	dataMap := make(map[string][]byte)
	oldData := h.Read(sid)
	if oldData != nil {
		json.Unmarshal(oldData, &dataMap)
		if value, ok := dataMap[name]; ok {
			dataBf := bytes.NewBuffer(value)
			dglog.Debugf("Session decoder:sid:%s, name:%s, data:%v \n", sid, name, value)
			dec := gob.NewDecoder(dataBf)
			dec.Decode(data)
		}
	}
}

func (manager *sessioneManager) delete(sid string, name string) {
	h := manager.getHandle()
	dataMap := make(map[string]interface{})
	oldData := h.Read(sid)
	if oldData != nil {
		json.Unmarshal(oldData, &dataMap)

		if value, ok := dataMap[name]; ok {
			dglog.Debugf("Delete sessionID:%s, name:%s, value:%s", sid, name, value)
			delete(dataMap, name)
			newData, err := json.Marshal(dataMap)
			if err != nil {
				dglog.Errorf("Can't marshal session[sid:%s, name:%s, dataMap:%v] to session ", sid, name, dataMap)
				return
			}

			h.Write(sid, newData)
		} else {
			dglog.Debugf("Could not found sessionID:%s, name:%s, value:%s", sid, name, value)
		}
	}
}

func (manager *sessioneManager) getHandle() SessionHandle {
	var h SessionHandle
	manager.m.Lock()
	h = manager.handle
	manager.m.Unlock()
	return h
}

func (manager *sessioneManager) setHandle(sh SessionHandle) {
	manager.m.Lock()
	manager.handle = sh
	manager.m.Unlock()
}

func SetSessionStore(sh SessionHandle) {
	manager.setHandle(sh)
}

func GetSession(sid string) *SessionContainer {
	manager.m.Lock()
	defer manager.m.Unlock()

	if sessionContainer, ok := manager.container[sid]; ok {

		return sessionContainer
	} else {

		sessionContainer := &SessionContainer{
			sid:    sid,
			m:      &sync.Mutex{},
			expire: time.Now().Unix() + int64(time.Second*3600),
		}
		manager.container[sid] = sessionContainer
		return sessionContainer
	}
	return nil
}

func GenerateSid() (string, error) {
	var sid string
	urandf, err := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	if err == nil {
		defer urandf.Close()
		data := make([]byte, 16)
		urandf.Read(data)
		sid = fmt.Sprintf("%x-%x-%x-%x-%x", data[:4], data[4:6], data[6:8], data[8:10], data[10:])
	} else {
		b := make([]byte, 16)
		n, err := rand.Read(b)
		if n != len(b) || err != nil {
			return "", errors.New("Could not generate sid")
		}
		sid = hex.EncodeToString(b)
		sid = fmt.Sprintf("%s-%s-%s-%s-%s", sid[:8], sid[8:12], sid[12:16], sid[16:20], sid[20:])
	}
	return sid, nil
}

func init() {
	manager = &sessioneManager{m: &sync.Mutex{}, container: make(map[string]*SessionContainer)}
	// 设置FileStoreEntity为默认的session存储驱动
	fs := handle.FileStoreEntity
	err := fs.Open()
	if err != nil {
		dglog.Errorf("Could not call FileStoreEntity open menthod:%v ", err)
	}
	manager.setHandle(fs)
}
