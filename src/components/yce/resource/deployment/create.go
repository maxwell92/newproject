package deployment

import (
	"api/controller"
	"fmt"
	//kubernetes "components/kubernetes/resource"
	//mysql "components/mysql/resource"

	"components/mysql"
	"components/yce/resource"
	"github.com/kataras/iris"
)

type CreateDeploymentController struct {
	controller.Handler
	Request  *resource.YceResource
	Response *controller.Response
}

func (cdc CreateDeploymentController) Url(url string)       {}
func (cdc CreateDeploymentController) Regist()              { iris.API(cdc.URL, cdc) }
func (cdc CreateDeploymentController) Method(method string) {}
func (cdc CreateDeploymentController) OK(message string)    { cdc.Write(message) }
func (cdc CreateDeploymentController) ERROR()               { cdc.Write("ERROR") }

func (cdc CreateDeploymentController) Post() {

	org := mysql.Organization().Query(cdc.Request.Metadata.OrgId)
	// org := mysql.Organization().Query().By(cdc.Request.Metadata.OrgId)

	/*
		for _, dcId := range cdc.Request.Metadata.DcIdList {
			dc := mysql.DataCenter().Query().By(dcId)
			apiServer := dc.Host + ":" + dc.Port
			kubernetes.DataCenter(apiServer).Namespace(org.Name).Deployment(cdc.Request.Spec).Create()

			mysql.Deployment(cdc.Metadata.userId, cdc.Metadata.DcIdList, mysql.ACTION_DEPLOY).Insert()
		}
	*/

	fmt.Println("this is createDeployment Controller")
	cdc.OK("OK")
}
