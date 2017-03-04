package deployment

import (
	controller "api/controller"
	"components/yce/resource"
)

type Deployment struct {
	resource.IResource
	Handlers []controller.IHandler
}

func New() *Deployment {
	return &Deployment{
		Handlers: make([]controller.IHandler, 0),
	}
}
