package users

import (
	controller "api/controller"
	"components/mysql/resource"
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

func Create() *CreateUserController {

}

func Check() *CheckUserController {

}
