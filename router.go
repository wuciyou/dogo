package dogo

import (
	"strings"
	"sync"
)

type routerContainer struct {
	ch     ContextHandle
	method HttpMethod
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

func (r *muxEntry) set(pattern string, method HttpMethod, ch ContextHandle) {
	r.mu.Lock()
	defer r.mu.Unlock()

	pattern = strings.TrimSpace(pattern)

	if _, ok := r.routerMap[pattern]; ok {
		DogoLog.Errorf("The pattern[%s] already exists in the Router", pattern)
		return
	}
	rc := &routerContainer{method: method, ch: ch}
	r.routerMap[pattern] = rc
}

func (r *muxEntry) Router(pattern string, ch ContextHandle) {
	r.set(pattern, "", ch)
}

func (r *muxEntry) GetRouter(pattern string, ch ContextHandle) {
	r.set(pattern, GET, ch)
}

func (r *muxEntry) PostRouter(pattern string, ch ContextHandle) {
	r.set(pattern, POST, ch)
}

func (r *muxEntry) PutRouter(pattern string, ch ContextHandle) {
	r.set(pattern, PUT, ch)
}

func (r *muxEntry) DeleteRouter(pattern string, ch ContextHandle) {
	r.set(pattern, DELETE, ch)
}

func Router(pattern string, ch ContextHandle) {
	router.set(pattern, "", ch)
}

func GetRouter(pattern string, ch ContextHandle) {
	router.set(pattern, GET, ch)
}

func PostRouter(pattern string, ch ContextHandle) {
	router.set(pattern, POST, ch)
}

func PutRouter(pattern string, ch ContextHandle) {
	router.set(pattern, PUT, ch)
}

func DeleteRouter(pattern string, ch ContextHandle) {
	router.set(pattern, DELETE, ch)
}
