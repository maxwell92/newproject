package users

import (
	"api/controller"
	"components/mysql/resource"
	"fmt"

	"github.com/kataras/iris"
)

type CheckUserController struct {
	controller.Handler
	Request  *CheckUserRequest
	Response *CheckUserResponse
}

func (cuc CheckUserController) Url(url string)       {}
func (cuc CheckUserController) Regist()              { iris.API(cuc.URL, cuc) }
func (cuc CheckUserController) Method(method string) {}
func (cuc CheckUserController) OK(message string)    { cuc.Write(message) }
func (cuc CheckUserController) ERROR()               { cuc.Write("ERROR") }

type CheckUserRequest struct {
	resource.MySQLResource
}

type CheckUserResponse struct {
	resource.MySQLResource
}

func (cuc CheckUserController) Post() {
	fmt.Println("this is checkUserController post")

	// session validation / authentication

	// business logic

	// write back result
	cuc.OK("OK")
}

func (cuc CheckUserController) Get() {
	fmt.Println("this is checkUserController get")
	cuc.ERROR()
}
