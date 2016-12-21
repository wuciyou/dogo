package hooks

import (
	"github.com/wuciyou/dogo/common"
	"sync"
)

type HookHander func(param ...interface{})

type hook struct {
	m    *sync.Mutex
	tags map[common.HookTagName][]HookHander
}

var Hook = &hook{m: &sync.Mutex{}, tags: make(map[common.HookTagName][]HookHander)}

func (h *hook) Add(name common.HookTagName, hander func(param ...interface{})) {
	// 加锁
	h.m.Lock()
	defer h.m.Unlock()

	h.tags[name] = append(h.tags[name], HookHander(hander))

}

func (h *hook) each(name common.HookTagName, f func(hander HookHander)) {

	var tags []HookHander

	h.m.Lock()

	tags = h.tags[name]

	h.m.Unlock()

	if len(tags) > 0 {
		for _, t := range tags {
			f(t)
		}
	}
}

func (h *hook) Listen(name common.HookTagName, param ...interface{}) {
	h.each(name, func(hander HookHander) {
		hander(param)
	})
}

func Add(name common.HookTagName, hander func(param ...interface{})) {
	Hook.Add(name, hander)
}

func Listen(name common.HookTagName, param ...interface{}) {
	Hook.Listen(name, param)
}
