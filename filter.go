package dogo

import (
	"path"
	"strings"
)

type filterFunc func(ctx *Context) bool

type filterChain struct {
	childs map[string]*filterChain
	fh     []filterFunc
}

var filter_entity = &filterChain{childs: make(map[string]*filterChain)}

func (f *filterChain) doFilter(request_url string, ctx *Context) bool {
	request_url = path.Clean(request_url)
	rules := strings.Split(request_url, "/")
	return f.do(rules, ctx)
}

func (f *filterChain) do(rules []string, ctx *Context) bool {

	if f.fh != nil {
		for _, fh := range f.fh {
			Dglog.Debugf("fh:%v", fh)
			if fh == nil {
				continue
			}
			is_continue := fh(ctx)
			if is_continue == false {
				return false
			}
		}
	}
	if len(rules) == 0 || len(f.childs) == 0 {
		return true
	}
	for index, rule := range rules {
		if rule == "" {
			continue
		}
		if child, exsit := f.childs[rule]; exsit {
			if len(rules) > index {
				return child.do(rules[index+1:], ctx)
			}
		}
		return true
	}
	return true
}

func (f *filterChain) AddFilter(rule string, filterHandle filterFunc) {
	rule = path.Clean(rule)
	Dglog.Debugf("add filter by rule:%s", rule)
	f.add(strings.Split(rule, "/"), filterHandle)
}

func (f *filterChain) add(rules []string, filterHandle filterFunc) {
	if len(rules) <= 0 {
		rules = append(rules, "root")
	}

	if f.childs == nil {
		f.childs = make(map[string]*filterChain)
	}

	for index, rule := range rules {
		if rule == "" {
			continue
		}
		if child_filter_chain, exist := f.childs[rule]; exist {
			child_filter_chain.add(rules[index+1:], filterHandle)
			return
		} else {
			if len(rules)-1 == index {
				f.fh = append(f.fh, filterHandle)
			} else {
				child_filter := &filterChain{}
				f.childs[rule] = child_filter
				child_filter.add(rules[index+1:], filterHandle)
			}
		}
		break

	}
}
