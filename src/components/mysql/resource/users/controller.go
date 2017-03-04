package users

import (
	"api/router"
)

func (u *User) AddControllers() *User {

	cuc := new(CheckUserController)
	cuc.URL = "/api/v1/user/check"
	cuc.METHOD = "POST"
	u.Handlers = append(u.Handlers, cuc)

	cpc := new(CheckPasswordController)
	cpc.URL = "/api/v2/user/password"
	cpc.METHOD = "POST" //TODO: for security concern, it's better GET checksum from server then verify
	u.Handlers = append(u.Handlers, cpc)

	//...

	for _, h := range u.Handlers {
		router.Instance().Add(h)
	}

	return u
}
