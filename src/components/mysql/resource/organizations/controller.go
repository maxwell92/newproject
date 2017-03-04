package organizations

import (
	"api/router"
)

func (o *Organization) AddControllers() *Organization {

	chc := new(CheckOrganizationController)
	chc.URL = "/api/v1/organization/check"
	chc.METHOD = "POST"
	o.Handlers = append(o.Handlers, chc)

	cdc := new(CreateOrganizationController)
	cdc.URL = "/api/v1/organization"
	cdc.METHOD = "POST"
	o.Handlers = append(o.Handlers, cdc)

	//...

	for _, h := range o.Handlers {
		router.Instance().Add(h)
	}

	return o
}
