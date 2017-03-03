package users

import (
	"fmt"
	"api/controller"
	"github.com/kataras/iris"
)


type CheckUserController struct {
	controller.Handler
}

func (cuc CheckUserController) Url(url string) {}
func (cuc CheckUserController) Regist() {
	iris.API(cuc.URL, cuc)
}

func (cuc CheckUserController) Method(method string)  {}
func (cuc CheckUserController) OK(message string) {}
func (cuc CheckUserController) ERROR() {}

func (cuc CheckUserController) Post() {
	fmt.Println("this is checkUserController")
}

func (cuc CheckUserController) Get() {
	fmt.Println("this is checkUserController")
}
