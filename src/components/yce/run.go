package yce

import (
	"api/router"
	user "components/mysql/resource/users"
	"github.com/kataras/iris"
)

func Run() {
	user.New().AddControllers()

	router.Instance().Regist()

	// iris.StaticServe()
	iris.Listen(":8080")
}
