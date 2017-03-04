package organizations

import (
	controller "api/controller"
	"components/mysql/resource"
	"reflect"
)

type Organization struct {
	Resource *resource.MySQLResource
	Handlers []controller.IHandler
}

func New() *Organization {
	return &Organization{
		Handlers: make([]controller.IHandler, 0),
	}
}

func (o *Organization) Insert() *CreateOrganizationController {
	return &CreateOrganizationController{}
}

func (o *Organization) Check() *CheckOrganizationController {
	return &CheckOrganizationController{}
}

func (o *Organization) by(column interface{}) *Organization {
	switch reflect.TypeOf(column) {
	case reflect.Int32: // always it's the id
		o = o.byId(column)
	case reflect.String: // always it's the name
		o = o.byName()
	case reflect.Interface: // always it's the map: like (name, version)
		o = o.byOther()
	}
	return o
}

func (o *Organization) byId(column interface{}) *Organization {
	// o.Resource.Metadata.Id = int(column)
}

func (o *Organization) byName() *Organization {

}

func (o *Organization) byOther() *Organization {

}

func (o *Organization) Query(column interface{}) *Organization {
	org := o.by(column)

	q := new(QueryOrganizationController)
	q.Metadata = org.Resource.Metadata
	q.Spec = org.Resource.Spec.(QuerySpec)

	q.Query(QUERY_BY_ID)

	o.Resource.Metadata = q.Metadata
	o.Resource.Spec = q.Spec

	return o

	/*
	q := &QueryOrganizationController{
		MySQLResource: &resource.MySQLResource{
			Metadata: org.Resource.Metadata,
		},
	}

	q := o.Resource.Query(QUERY_BY_ID).(*QueryOrganizationController)
	*/
}

/*
func (o *Organization) SoftDelete() *DeleteOrganizationController {

}
*/
