package organizations

import (
	"components/mysql/resource"
	"reflect"

	"github.com/maxwell92/newutils/mysql"
)

type QueryOrganizationController struct {
	resource.MySQLResource
}

func (q *QueryOrganizationController) Insert(SQL string) *QueryOrganizationController     {}
func (q *QueryOrganizationController) Query(SQL string) *QueryOrganizationController      {
	db := mysql.MysqlInstance().Conn()

	stmt, err := db.Prepare(SQL)
	if err != nil {

	}

	err1 := stmt.QueryRow(q.Metadata.Id, VALID).Scan(q.Metadata.Id, q.Metadata.Name, )
	if err1 != nil {

	}



}

func (q *QueryOrganizationController) SoftDelete(SQL string) *QueryOrganizationController {}

