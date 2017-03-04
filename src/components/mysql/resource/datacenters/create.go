package datacenters

import (
	"api/controller"
	"components/mysql/resource"

	"github.com/kataras/iris"
)

type CreateDatacenterController struct {
	controller.Handler
	Request  *CreateDatacenterRequest
	Response *CreateDatacenterResponse
}

func (cdc CreateDatacenterController) Url(url string)       {}
func (cdc CreateDatacenterController) Regist()              { iris.API(cdc.URL, cdc) }
func (cdc CreateDatacenterController) Method(method string) {}
func (cdc CreateDatacenterController) OK(message string)    { cdc.Write(message) }
func (cdc CreateDatacenterController) ERROR()               { cdc.Write("ERROR") }

type CreateDatacenterRequest struct {
	resource.MySQLResource
}

type CreateDatacenterResponse struct {
	resource.MySQLResource
}
