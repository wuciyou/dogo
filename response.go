package dogo

import (
	"bytes"
)

type Response struct {
	data *bytes.Buffer // 输出内容
	code int           // 输出状态码
}

func (w *Response) Date(data []byte) {

}

func (w *Response) Send() {

}
