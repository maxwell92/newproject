package yce

import (
	"api/router"
)

func Run() {
	// router.Regist()
	router.New().Regist()
	// iris.Listen()
}
