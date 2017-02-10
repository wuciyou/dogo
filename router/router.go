package router

import (
	"errors"
	"fmt"
	"github.com/wuciyou/dogo/common"
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
	"net/http"
	"strings"
	"sync"
)

type routerContainer struct {
	ch     context.ContextHandle
	method common.HttpMethod
	host   string
}

type muxEntry struct {
	routerMap map[string]*routerContainer
	mu        sync.RWMutex
}

var routerEnttry = &muxEntry{routerMap: make(map[string]*routerContainer)}

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

// admin@index/index:get(b[])|post
//
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
	routerEnttry.set(pattern, "", ch)
}

func GetRouter(pattern string, ch context.ContextHandle) {
	routerEnttry.GetRouter(pattern, ch)
}

func PostRouter(pattern string, ch context.ContextHandle) {
	routerEnttry.PostRouter(pattern, ch)
}

func PutRouter(pattern string, ch context.ContextHandle) {
	routerEnttry.PutRouter(pattern, ch)
}

func DeleteRouter(pattern string, ch context.ContextHandle) {
	routerEnttry.DeleteRouter(pattern, ch)
}

func Match(request *http.Request) (context.ContextHandle, string, error) {

	routerContainer, pattern := routerEnttry.match(request.URL.Path)
	if routerContainer == nil {
		return nil, pattern, errors.New("404 page not found")
	}

	if routerContainer.method != "" && string(routerContainer.method) != request.Method {
		// http.NotFound(response, request)

		return nil, pattern, errors.New(fmt.Sprintf("Request method[%s] must be %s,url:%s", request.Method, routerContainer.method, request.URL.Path))
	}

	return routerContainer.ch, pattern, nil

}
