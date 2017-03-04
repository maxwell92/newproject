package deployment

import (
	"api/router"
)

func (d *Deployment) AddControllers() *Deployment {
	chc := new(CheckDeploymentController)
	chc.URL = "/api/v1/deployment/check"
	chc.METHOD = "POST"
	d.Handlers = append(d.Handlers, chc)

	cdc := new(CreateDeploymentController)
	cdc.URL = "/api/v1/deployment"
	cdc.METHOD = "POST"
	d.Handlers = append(d.Handlers, cdc)

	//...

	for _, h := range d.Handlers {
		router.Instance().Add(h)
	}

	return d
}
