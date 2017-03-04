package users

import (
	"api/router"
)

func (u *User) AddControllers() *User {

	/*
		cuc := new(CheckUserController)
		cuc.URL = "/api/v1/user/check"
		//cuc.URL = "/"
		cuc.METHOD = "POST"
	*/

	cuc := &CheckUserController{
		URL:    "/api/v2/user/check",
		METHOD: "POST",
	}
	u.Handlers = append(u.Handlers, cuc)

	cpc := &CheckPasswordController{
		URL:    "/api/v2/user/password",
		METHOD: "POST", //TODO: for security concern, it's better GET checksum from server then verify
	}
	u.Handlers = append(u.Handlers, cpc)

	//...

	for _, h := range u.Handlers {
		router.Instance().Add(h)
	}

	return u
}
