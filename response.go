package dogo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"
)

type Response struct {
	http.ResponseWriter
	data       *bytes.Buffer
	assignData map[string]interface{}
}

func initResponse(w http.ResponseWriter) Response {
	response := Response{
		ResponseWriter: w,
		data:           &bytes.Buffer{},
		assignData:     make(map[string]interface{}),
	}
	return response
}

func (w Response) Write(data []byte) (int, error) {
	return w.data.Write(data)
}

func (w Response) Assign(name string, data interface{}) {
	w.assignData[name] = data
}

func (w Response) Display(templateFileName ...string) {
	if len(templateFileName) <= 0 {
		Dglog.Error("Must need templage file name")
		return
	}
	template := NewTemplate(templateFileName[0])
	template.assignData = w.assignData
	template.Display(w)
	w.Send()
}

/**
 * 输出json格式数据
 *
 */
func (w Response) Json(jsonData interface{}) {
	if data, err := json.Marshal(jsonData); err == nil {
		w.data.Write(data)
		w.SetContentType("application/json")
		w.Send()
	} else {
		Dglog.Errorf("Can't marshal json data for %+v", jsonData)
	}
}

/**
 * 输出xml格式数据
 *
 */
func (w Response) Xml(xmlData interface{}) {
	if data, err := xml.Marshal(xmlData); err == nil {
		w.data.Write(data)
		w.SetContentType("application/xml")
		w.Send()
	} else {
		Dglog.Errorf("Can't marshal xml data for %+v", xmlData)
	}
}

/**
 * 输出到浏览器
 *
 */
func (w Response) Send() {
	w.data.WriteTo(w.ResponseWriter)
}

func (w Response) SetContentType(contentType string, charset ...string) {

	if contentType == "" {
		contentType = "text/html"
	}
	if len(charset) == 0 {
		charset = []string{"utf-8"}
	} else if charset[0] == "" {
		charset[0] = "utf-8"
	}
	w.ResponseWriter.Header().Add("Content-Type", contentType+"; charset="+charset[0])
}

func (w Response) Redirect(remoteUrl *url.URL) {
}
