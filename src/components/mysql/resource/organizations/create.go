package organizations

import (
	"api/controller"
	"components/mysql/resource"
	"github.com/kataras/iris"
)

type CreateOrganizationController struct {
	controller.Handler
	Request  *CreateOrganizationRequest
	Response *CreateOrganizationResponse
}

func (coc CreateOrganizationController) Url(url string)       {}
func (coc CreateOrganizationController) Regist()              { iris.API(coc.URL, coc) }
func (coc CreateOrganizationController) Method(method string) {}
func (coc CreateOrganizationController) OK(message string)    { coc.Write(message) }
func (coc CreateOrganizationController) ERROR()               { coc.Write("ERROR") }

type CreateOrganizationRequest struct {
	resource.MySQLResource
}

type CreateOrganizationResponse struct {
	resource.MySQLResource
}

func (coc CreateOrganizationController) Post() {
	//mysql.Organization().Insert()

}
