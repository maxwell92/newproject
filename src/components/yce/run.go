package yce

import (
	"api/router"
	"components/mysql"
	"github.com/kataras/iris"
)

func Run() {
	mysql.AddControllers()
	AddControllers()

	router.Instance().Regist()

	// iris.StaticServe()
	iris.Listen(":8080")
}
