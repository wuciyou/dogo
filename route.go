package dogo

import (
	"net/http"
	"path"
	"strings"
)

type route struct {
	rules map[string]*route
	h     handle
	hooks map[string]func(string)
}

var route_entity = &route{rules: make(map[string]*route), hooks: make(map[string]func(string))}

type handle func(c *Context)

func (r *route) init(method string) {

	if _, ok := r.rules[method]; !ok {
		r.rules[method] = &route{rules: make(map[string]*route)}
	}
}

func (r *route) Filter(f func(), filterHandle filterFunc) {
	r.hooks["filter"] = func(rule string) {
		filter_entity.AddFilter(rule, filterHandle)
	}
	f()
	r.hooks["filter"] = nil

}
func (r *route) doHooks(name string, data string) {
	if r.hooks == nil {
		return
	}
	if hook, exist := r.hooks[name]; exist && hook != nil {
		hook(data)
	}
}
func (r *route) Any(rule string, h handle) {
	r.Get(rule, h)
	r.Post(rule, h)
}

func (r *route) Get(rule string, h handle) {
	r.init("GET")
	r.doHooks("filter", rule)
	(r.rules["GET"]).addRoute(strings.Split(rule, "/"), h)
}

func (r *route) Post(rule string, h handle) {
	r.init("POST")
	r.doHooks("filter", rule)
	(r.rules["POST"]).addRoute(strings.Split(rule, "/"), h)
}

func (r *route) addRoute(rule []string, h handle) {
	Dglog.Debugf("add route %s ", rule)

	if len(rule) > 1 && rule[0] == "" {
		rule = rule[1:]
	} else {
		rule = rule
	}

	var nextRoute *route
	if nextRouteTemp, ok := r.rules[rule[0]]; ok {
		nextRoute = nextRouteTemp
	} else {
		nextRoute = &route{rules: make(map[string]*route)}
	}

	if len(rule) == 1 {
		nextRoute.h = h
	} else {
		nextRoute.addRoute(rule[1:], h)
	}
	r.rules[rule[0]] = nextRoute
}

func (r *route) checkRoute(request *http.Request) handle {
	method := request.Method
	// 将后缀去掉
	requestUri := strings.TrimSuffix(request.URL.Path, path.Ext(request.URL.Path))
	url := strings.Split(requestUri, "/")
	h := r.checkMethod(method, url)
	return h
}

func (r *route) checkMethod(method string, url []string) handle {
	return (r.rules[strings.ToUpper(method)]).checkUrl(url)
}

func (r *route) checkUrl(url []string) (h handle) {
	if len(url) <= 0 {
		return r.h
	} else if len(url) > 1 && url[0] == "" {
		url = url[1:]
	}

	if nextRule, exist := r.rules[url[0]]; exist {
		h = nextRule.checkUrl(url[1:])
	}
	if h == nil {
		Dglog.Debugf("进行模糊匹配")
		for k, v := range r.rules {
			ok, err := path.Match(k, url[0])
			if ok && err == nil {
				h = v.checkUrl(url[1:])
				if h != nil {
					return
				}
				break
			}
		}
	}
	return
}
