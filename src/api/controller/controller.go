package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
)

type IHandler interface {
	Url(url string)
	Method(method string)
	Regist()

	// write ok message back to frontend
	OK(message string)
	// write error message back to frontend
	ERROR()
}

type Handler struct {
	*iris.Context
	URL    string
	METHOD string
}

func (h Handler) Regist() {
	iris.API(h.URL, h)
}

func (h Handler) Url(url string) {
	h.URL = url
}

func (h Handler) Method(method string) {
	h.METHOD = method
}

func (h Handler) OK(message string) {
	// h.Write(message)
	resp := &Response{
		Code: 0,
		Msg:  message,
		Data: "",
	}

	h.Write(resp.Encode())
}

func (h Handler) ERROR() {
	// h.Write("ERROR")
}

type Handlers []Handler

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (resp Response) Encode() string {
	data, err := json.Marshal(resp)
	if err != nil {

	}

	return string(data)
}
