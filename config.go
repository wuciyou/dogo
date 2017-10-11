package dogo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var configMutex = &sync.RWMutex{}
var defaultConfig = dogoConfig{
	// 服务相关配置
	"SERVER": dogoConfig{
		// 服务器监听端口
		"WEB_PORT": 8108,
		"WEB_IP":   "",
		// web服务器名称（系统默认会添加到response headers返回头信息）
		"WEB_NAME": "DoGoServerv/1.0",
	},
}

type dogoConfig map[string]interface{}

func (this dogoConfig) Set(value interface{}, names ...string) error {
	configMutex.Lock()
	defer configMutex.Unlock()
	if names == nil {
		return errors.New("Name cant't null")
	}
	if len(names) == 1 {
		names = strings.Split(names[0], ".")
	}
	var temp = this
	namesLen := len(names)
	for i, name := range names {
		if i == namesLen-1 {

			Debugf("old:%v, temp[%s]= %v", temp, name, value)
			temp[name] = value
			Debugf("new:%v, temp[%s]= %v", temp, name, value)
		} else {
			if v, ok := temp[name]; ok {
				switch v.(type) {
				case dogoConfig:
					temp = v.(dogoConfig)
				default:
					return errors.New(fmt.Sprintf("conf[%v] must is dogoConfig type", names))
				}
			} else {
				temp[name] = dogoConfig{}
			}
		}
	}
	return nil
}

func (this dogoConfig) Get(names ...string) interface{} {
	configMutex.Lock()
	defer configMutex.Unlock()

	if !this.checkName(names...) {
		Debugf("get config param:%v, not found the keys", names)
		return nil
	}

	if len(names) == 1 {
		Debugf("get config param:%v, value:%v", names, this[names[0]])
		return this[names[0]]
	}

	var temp = this

	for _, v := range names {

		switch temp[v].(type) {
		case dogoConfig:
			temp = temp[v].(dogoConfig)
		case interface{}:
			Debugf("get config param:%v, value:%v", names, temp[v])
			return temp[v]
		}
	}
	Debugf("get config param:%v, not found the keys", names)
	return nil
}

func (this dogoConfig) checkName(name ...string) bool {
	if name == nil {
		return false
	} else if len(name) == 0 {
		return false
	}
	return true
}

func (this dogoConfig) String(name ...string) string {
	var values interface{}
	if len(name) > 1 {
		values = this.Get(name...)
	} else {
		values = this.Get((strings.Split(name[0], "."))...)
	}

	if values == nil {
		return ""
	}
	switch values.(type) {
	case string:
		return values.(string)
	case int, byte, int8, int16, int32, int64, uint, uint16, uint32, uint64:
		return fmt.Sprintf("%d", values)
	}
	return ""
}

func (this dogoConfig) Int(name ...string) int {
	var values interface{}
	if len(name) > 1 {
		values = this.Get(name...)
	} else {
		values = this.Get((strings.Split(name[0], "."))...)
	}
	if values == nil {
		return 0
	}
	switch values.(type) {
	case int:
		return values.(int)
	}
	return 0
}

func (this dogoConfig) copy() dogoConfig {
	confJson, _ := json.Marshal(this)
	var tempConfg dogoConfig
	json.Unmarshal(confJson, &tempConfg)
	return tempConfg
}

func setConf(value interface{}, names ...string) error {
	return defaultConfig.Set(value, names...)
}

func getConfInt(name string) int {
	return defaultConfig.Int(name)
}

func getConfString(name string) string {
	return defaultConfig.String(name)
}
