package dogo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
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
	Pattern  string
	Suffix   string
}

func (c *Context) AddHeader(name string, value string) {
	c.response.rw.Header().Add(name, value)
}

func (c *Context) Header() http.Header {
	return c.response.rw.Header()
}

func (c *Context) Write(data []byte) (int, error) {
	return c.response.writeBuf.Write(data)
}

func (c *Context) WriteString(data string) (n int, err error) {
	return c.response.writeBuf.WriteString(data)
}

func (c *Context) WriteRune(data rune) (n int, err error) {
	return c.response.writeBuf.WriteRune(data)
}

func (c *Context) WriteByte(data byte) error {
	return c.response.writeBuf.WriteByte(data)
}

func (c *Context) WriteHeader(i int) {
	c.response.rw.WriteHeader(i)
}

func (c *Context) AjaxReturn(data interface{}, format ...string) {
	var ajaxReturnRormat string

	if len(format) > 0 {
		ajaxReturnRormat = format[0]
	} else if c.Suffix != "" {
		ajaxReturnRormat = c.Suffix
	} else {
		ajaxReturnRormat = RunTimeConfig.ajaxReturnRormat
	}

	switch strings.ToUpper(ajaxReturnRormat) {

	case "JSON":
		dataJson, err := json.Marshal(data)
		if err != nil {
			DogoLog.Errorf("AjaxReturn json marshal fail:%v \n", err)
			return
		}
		c.AddHeader("Content-Type", "application/json; charset=utf-8")
		c.response.writeBuf.Write(dataJson)
	case "XML":
		dataXml, err := xml.Marshal(data)
		if err != nil {
			DogoLog.Errorf("AjaxReturn xml marshal fail:%v \n", err)
			return
		}
		c.AddHeader("Content-Type", "text/xml; charset=utf-8")
		c.response.writeBuf.Write(dataXml)
	}
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

func (c *Context) parse(response http.ResponseWriter, request *http.Request) {

	c.response = &dogoResponse{rw: response, writeBuf: bytes.NewBuffer(make([]byte, 0))}
	c.Request = request

	c.Request.URL.Path = filepath.Clean(c.Request.URL.Path)
	suffixPoint := strings.IndexAny(c.Request.URL.Path, ".")

	if suffixPoint >= 0 {
		c.Pattern = c.Request.URL.Path[:suffixPoint]
		c.Suffix = c.Request.URL.Path[suffixPoint+1:]
	} else {
		c.Pattern = c.Request.URL.Path
	}

}
