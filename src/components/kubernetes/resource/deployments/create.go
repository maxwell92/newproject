package deployments

import (
	"components/mysql/resource/users"
)

func Post() {
	// users.Controller(users.CreateUserController).Insert()
	users.Check().Request
}
