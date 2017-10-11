package dogo

import (
	"testing"
)

func wuciyou(ctx *Context) bool {
	Dglog.Debug("进入到wuciyou")
	if ctx.runtimeContainer == nil {
		ctx.runtimeContainer = make(map[string]interface{})
		ctx.runtimeContainer["wuciyou"] = 12
	} else {
		ctx.runtimeContainer["iiiwuciyou"] = 12
	}
	Dglog.Debug("runtimeContainer:%+v", ctx.runtimeContainer)
	return false
}

func TestAdd(t *testing.T) {

	filter := &FilterChain{}
	filter.AddFilter("/hello/wuciyou", wuciyou)
	filter.AddFilter("/whello/ewuciyou", wuciyou)
	filter.AddFilter("/hello/", wuciyou)
	ctx := &Context{}
	is_continue := filter.doFilter("/hello/wuciyou", ctx)
	t.Logf("filter :%+v is_continue:%b ", filter, is_continue)
}
