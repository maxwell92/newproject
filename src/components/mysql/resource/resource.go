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

// MySQLResource is the metadata of every MySQL resource object
type MySQLResource struct {
	CreatedAt  string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
	ModifiedOp int32  `json:"modifiedOp"`
	Comment    string `json:"comment"`
}

/*
type MySQLResource Metadata

type MySQLResource struct {
	*Metadata `json:"metadata"`
}
*/

func (mr *MySQLResource) Insert() MySQLResource {
	return mr
}

func (mr *MySQLResource) Query() MySQLResource {

	return mr
}

func (mr *MySQLResource) SoftDelete() MySQLResource {

	return mr
}

// by the specific column
func (mr *MySQLResource) By() MySQLResource {

	return mr
}

func test() {
	//	m := new(MySQLResource)
}
