package users

import (
	"components/mysql/resource"
	controller "api/controller"
)

type User struct {
	resource.MySQLResource
	Handlers []controller.IHandler
}

func New() *User {
	return &User{
		Handlers: make([]controller.IHandler, 0),
	}
}