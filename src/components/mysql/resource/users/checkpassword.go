package users

import (
	"api/controller"
	"fmt"
	"github.com/kataras/iris"
)

type CheckPasswordController struct {
	controller.Handler
}

func (cpc CheckPasswordController) Url(url string) {}
func (cpc CheckPasswordController) Regist() {
	iris.API(cpc.URL, cpc)
}

func (cpc CheckPasswordController) Method(method string) {}
func (cpc CheckPasswordController) OK(message string)    { cpc.Write(message) }
func (cpc CheckPasswordController) ERROR()               { cpc.Write("ERROR") }

func (cpc CheckPasswordController) Post() {
	fmt.Println("this is checkPasswordController post")
	cpc.OK("Post")
}

func (cpc CheckPasswordController) Get() {
	fmt.Println("this is checkPasswordController get")
	cpc.ERROR()
}
