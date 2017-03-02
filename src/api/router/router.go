package router

import (
	"github.com/kataras/iris"
	"api/controller"
)

type Router struct {
	controller.Controllers
}

func New() *Router {
	r := new(Router)

}

func (r *Router) init() {
	for _, c := range r.Controllers {

	}
}

func (r *Router) Regist() {

}






