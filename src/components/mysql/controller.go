package mysql

import (
	"components/mysql/resource/datacenters"
	"components/mysql/resource/organizations"
	"components/mysql/resource/users"
)

func AddControllers() {
	users.New().AddControllers()
	datacenters.New().AddControllers()
	organizations.New().AddControllers()
}

func Organization() *organizations.Organization {
	return &organizations.Organization{}
}

/*
type IOrganization interface {
	*QueryOrganizationInterface
}

type QueryOrganizationInterface interface {
	Query()
}

func Organization() IOrganization {

}
*/

func Datacenter() {

}
