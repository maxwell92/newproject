package resource

import (
	"components/yce/resource"
)

type IMySQLResource interface {
	// Insert is insert on duplicate update
	Insert(string) IMySQLResource
	Query(string) IMySQLResource
	SoftDelete(string) IMySQLResource
}

type MySQLResource struct {
	Metadata *MySQLMetadata
	Spec     MySQLSpecInterface
}

func (mr *MySQLResource) exec(SQL string) IMySQLResource {

}

func (mr *MySQLResource) Query(SQL string) IMySQLResource {
	obj := mr.exec(SQL)
	return obj
}
/*
func (mr *MySQLResource) Insert(SQL string) MySQLResource {
	return *mr
}

func (mr *MySQLResource) SoftDelete(SQL string) MySQLResource {

	return *mr
}

// by the specific column
func (mr *MySQLResource) By() MySQLResource {

	return *mr
}
*/

func test() {
	//	m := new(MySQLResource)
}

type MySQLMetadata struct {
	Id         int32  `json:"id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
	ModifiedOp int32  `json:"modifiedOp"`
	Comment    string `json:"comment"`
}

type MySQLSpecInterface interface {
	Encode() string
	Decode([]byte)
}
