package router

import (
	"api/controller"
	"sync"
)

var once sync.Once
var r *Router
type Router struct {
	Handlers []controller.IHandler
}

func Instance() *Router {
	once.Do(func() {
		r = new(Router)
		r.Handlers = make([]controller.IHandler, 0)
	})
	return r
}

func (r *Router) Add(c controller.IHandler) {
	r.Handlers = append(r.Handlers, c)
}

func (r *Router) Regist() {
	for _, h := range r.Handlers {
		h.Regist()
	}
}
