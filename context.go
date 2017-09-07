package dogo

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Context struct {
	R        *Request
	W        Response
	form     url.Values
	postForm url.Values
	cookie   map[string]string
	// 自定义数据
	Container map[string]interface{}
	// 系统运行数据
	runtimeContainer map[string]interface{}
}

func InitContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{R: &Request{Request: r}, W: initResponse(w)}
	// 解析请求参数
	ctx.R.ParseForm()
	ctx.form = ctx.R.Form
	ctx.postForm = ctx.R.PostForm
	ctx.Container = make(map[string]interface{})
	ctx.runtimeContainer = make(map[string]interface{})
	return ctx
}

/**
 * 判断是否是post请求
 *
 */
func (ctx *Context) IsPost() bool {
	return ctx.R.Method == "POST"
}

/**
 * 判断是否是get请求
 *
 */
func (ctx *Context) IsGet() bool {
	return ctx.R.Method == "GET"
}

/**
 * 判断是否是put请求
 *
 */
func (ctx *Context) IsPut() bool {
	return ctx.R.Method == "PUT"
}

/**
 * 判断是否是ajax请求
 *
 */
func (ctx *Context) IsAjax() bool {
	return strings.ToLower(ctx.R.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

/**
 * 获取用户的ip和端口号
 *
 */
func (ctx *Context) ClientIp() (ip string, port int) {
	addrs := strings.Split(ctx.R.RemoteAddr, ":")
	ip = addrs[0]
	port, _ = strconv.Atoi(addrs[1])
	return
}

/**
 * 获取cookie值
 *
 */
func (ctx *Context) GetCookie(name string) string {
	if cookie, err := ctx.R.Cookie(name); err == nil && cookie != nil {
		return cookie.Value
	} else {
		Dglog.Warningf("Can't found cookie for name %s", name)
	}
	return ""
}

/**
 * 设置cookie值
 *
 */
func (ctx *Context) SetCookie(name string, value string) {
	cookie := &http.Cookie{Name: name, Value: value}
	http.SetCookie(ctx.W.ResponseWriter, cookie)
}

/**
 * 获取Get请求参数
 *
 */
func (ctx *Context) Get(name string) string {
	return ctx.form.Get(name)
}

func (ctx *Context) Gets(name string) []string {
	return ctx.form[name]
}

/**
 * 获取Post请求参数
 *
 */
func (ctx *Context) Post(name string) string {
	return ctx.postForm.Get(name)
}

func (ctx *Context) Posts(name string) []string {
	return ctx.postForm[name]
}

func (ctx *Context) File(name string) {
	//f, fstat, err := ctx.R.FormFile(name)
}
