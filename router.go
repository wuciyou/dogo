package dogo

import (
	"github.com/wuciyou/dogo/common"
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
	"strings"
	"sync"
	//
)

type routerContainer struct {
	ch     context.ContextHandle
	method common.HttpMethod
}

type muxEntry struct {
	routerMap map[string]*routerContainer
	mu        sync.RWMutex
}

var router = &muxEntry{routerMap: make(map[string]*routerContainer)}

func (r *muxEntry) match(path string) (*routerContainer, string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 匹配全路径
	if rc, ok := r.routerMap[path]; ok {
		return rc, path
	}

	suffixPoint := strings.IndexAny(path, ".")
	if suffixPoint >= 0 {
		path = path[:suffixPoint]
		if rc, ok := r.routerMap[path]; ok {
			return rc, path
		}
	}

	return nil, path
}

func (r *muxEntry) set(pattern string, method common.HttpMethod, ch context.ContextHandle) {
	r.mu.Lock()
	defer r.mu.Unlock()

	pattern = strings.TrimSpace(pattern)

	if _, ok := r.routerMap[pattern]; ok {
		dglog.Errorf("The pattern[%s] already exists in the Router", pattern)
		return
	}
	rc := &routerContainer{method: method, ch: ch}
	r.routerMap[pattern] = rc
}

func (r *muxEntry) Router(pattern string, ch context.ContextHandle) {
	r.set(pattern, "", ch)
}

func (r *muxEntry) GetRouter(pattern string, ch context.ContextHandle) {
	r.set(pattern, common.GET, ch)
}

func (r *muxEntry) PostRouter(pattern string, ch context.ContextHandle) {
	r.set(pattern, common.POST, ch)
}

func (r *muxEntry) PutRouter(pattern string, ch context.ContextHandle) {
	r.set(pattern, common.PUT, ch)
}

func (r *muxEntry) DeleteRouter(pattern string, ch context.ContextHandle) {
	r.set(pattern, common.DELETE, ch)
}

func Router(pattern string, ch context.ContextHandle) {
	router.set(pattern, "", ch)
}

func GetRouter(pattern string, ch context.ContextHandle) {
	router.GetRouter(pattern, ch)
}

func PostRouter(pattern string, ch context.ContextHandle) {
	router.PostRouter(pattern, ch)
}

func PutRouter(pattern string, ch context.ContextHandle) {
	router.PutRouter(pattern, ch)
}

func DeleteRouter(pattern string, ch context.ContextHandle) {
	router.DeleteRouter(pattern, ch)
}
