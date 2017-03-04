package deployment

import (
	"api/controller"
	"fmt"

	"github.com/kataras/iris"
)

type CheckDeploymentController struct {
	controller.Handler
}

func (cdc CheckDeploymentController) Url(url string)       {}
func (cdc CheckDeploymentController) Regist()              { iris.API(cdc.URL, cdc) }
func (cdc CheckDeploymentController) Method(method string) {}
func (cdc CheckDeploymentController) OK(message string)    { cdc.Write(message) }
func (cdc CheckDeploymentController) ERROR()               { cdc.Write("ERROR") }

func (cdc CheckDeploymentController) Post() {
	fmt.Println("this is checkDeployment Controller")
	cdc.OK("OK")
}
