package controller

import (
	"github.com/kataras/iris"
)

type IController interface {
	URL() IController
	Method() IController

	// write ok message back to frontend
	OK(string)
	// write error message back to frontend
	ERROR()
}

type Controller struct {
	iris.Context
}

func (c *Controller) URL() IController {

}

func (c *Controller) Method() IController {

}

func (c *Controller) OK(message string) {

}

func (c *Controller) ERROR() {

}

type Controllers []Controller
