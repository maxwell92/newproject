package datacenters

import (
	"api/router"
)

func (d *Datacenter) AddControllers() *Datacenter {

	chc := new(CheckDatacenterController)
	chc.URL = "/api/v1/datacenter/check"
	chc.METHOD = "POST"
	d.Handlers = append(d.Handlers, chc)

	cdc := new(CreateDatacenterController)
	cdc.URL = "/api/v1/datacenter"
	cdc.METHOD = "POST"
	d.Handlers = append(d.Handlers, cdc)

	//...

	for _, h := range d.Handlers {
		router.Instance().Add(h)
	}

	return d
}
