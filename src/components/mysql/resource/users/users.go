package users

import (
	"components/mysql/resource"
)

type Users struct {
	resource.MySQLResource
}

func New() {
	u := new(Users)
	u
}