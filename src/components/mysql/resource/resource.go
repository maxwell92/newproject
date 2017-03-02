package resource

import (
	"components/yce/resource"
)

type IMySQLResource interface {
	resource.IResource
	// Insert is insert on duplicate update
	Insert() IMySQLResource
	Query() IMySQLResource
	SoftDelete() IMySQLResource
}

type MySQLResource struct {
	CreatedAt string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
	ModifiedOp int32 `json:"modifiedOp"`
	Comment string `json:"comment"`
}

/*
type MySQLResource Metadata

type MySQLResource struct {
	*Metadata `json:"metadata"`
}

func (mr *MySQLResource) Insert() MySQLResource {

}

func (mr *MySQLResource) Query() MySQLResource {

}

func (mr *MySQLResource) SoftDelete() MySQLResource {

}

func (mr *MySQLResource) By() MySQLResource {

}

*/


func test() {
	m := new(MySQLResource)
}

