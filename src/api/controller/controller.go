package controller

import (
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
	URL string
	METHOD string
}

func (h Handler) Regist() {
	iris.API(h.URL, h)
}

func (h Handler) Url(url string) {
	h.URL = url
}

func (h Handler) GetUrl() string {
	return h.URL
}

func (h Handler) Method(method string) {
	h.METHOD = method
}

func (h Handler) OK(message string) {

}

func (h Handler) ERROR() {

}

type Handlers []Handler
