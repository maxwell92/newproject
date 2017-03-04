package datacenters

import (
	"api/controller"
	"components/mysql/resource"
	"fmt"

	"github.com/kataras/iris"
)

type CheckDatacenterController struct {
	controller.Handler
	Request  *CheckDatacenterRequest
	Response *CheckDatacenterResponse
}

func (cdc CheckDatacenterController) Url(url string)       {}
func (cdc CheckDatacenterController) Regist()              { iris.API(cdc.URL, cdc) }
func (cdc CheckDatacenterController) Method(method string) {}
func (cdc CheckDatacenterController) OK(message string)    { cdc.Write(message) }
func (cdc CheckDatacenterController) ERROR()               { cdc.Write("ERROR") }

type CheckDatacenterRequest struct {
	resource.MySQLResource
}

type CheckDatacenterResponse struct {
	resource.MySQLResource
}
