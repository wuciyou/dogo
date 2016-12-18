package dogo

import (
	"bytes"
	"html/template"
	"net/http"
)

type ContextHandle func(c *Context)

type dogoResponse struct {
	rw         http.ResponseWriter
	t          *template.Template
	writeBuf   *bytes.Buffer
	assignData []interface{}
}

type Context struct {
	response *dogoResponse
	Request  *http.Request
}

func (c *Context) parse(response http.ResponseWriter, request *http.Request) {

	wbuf := bytes.NewBuffer(make([]byte, 0))

	c.response = &dogoResponse{rw: response, writeBuf: wbuf}
	c.Request = request
}

func (c *Context) SetHeader(name string, value string) {
	c.response.rw.Header().Add(name, value)
	// DogoLog.Debugf("Add header to response ,key:%s ,value:%s, c.response.rw.Header():%+v", name, value, c.response.rw.Header())
}

func (c *Context) Header() http.Header {
	return c.response.rw.Header()
}

func (c *Context) Write(data []byte) (int, error) {
	return c.response.writeBuf.Write(data)
}

func (c *Context) WriteHeader(i int) {
	c.response.rw.WriteHeader(i)
}

func (c *Context) Display(templateFile ...string) {
	t, err := template.ParseFiles(templateFile...)
	if err != nil {
		DogoLog.Warningf("ParseFiles fail:%s", err)
		return
	}

	for _, v := range c.response.assignData {
		t.Execute(c.response.writeBuf, v)
	}

}

func (c *Context) Assign(data interface{}) {
	c.response.assignData = append(c.response.assignData, data)
}
