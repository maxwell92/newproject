package datacenters

import (
	controller "api/controller"
	"components/mysql/resource"
)

type Datacenter struct {
	resource.MySQLResource
	Handlers []controller.IHandler
}

func New() *Datacenter {
	return &Datacenter{
		Handlers: make([]controller.IHandler, 0),
	}
}

func Create() *CreateDatacenterController {
	return &CreateDatacenterController{}
}

func Check() *CheckDatacenterController {
	return &CheckDatacenterController{}
}
