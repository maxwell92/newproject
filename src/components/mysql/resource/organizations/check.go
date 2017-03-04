package organizations

import (
	"api/controller"
	"components/mysql/resource"

	"github.com/kataras/iris"
)

type CheckOrganizationController struct {
	controller.Handler
	Request  *CheckOrganizationRequest
	Response *CheckOrganizationResponse
}

func (coc CheckOrganizationController) Url(url string)       {}
func (coc CheckOrganizationController) Regist()              { iris.API(coc.URL, coc) }
func (coc CheckOrganizationController) Method(method string) {}
func (coc CheckOrganizationController) OK(message string)    { coc.Write(message) }
func (coc CheckOrganizationController) ERROR()               { coc.Write("ERROR") }

type CheckOrganizationRequest struct {
	resource.MySQLResource
}

type CheckOrganizationResponse struct {
	resource.MySQLResource
}
