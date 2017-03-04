package organizations

const (
	VALID   = 1
	INVALID = 2

	QUERY_BY_ID = "SELECT id, name, dcIdList, cpuQuota, memQuota, budget, balance, status, createdAt, modifiedAt, modifiedOp, comment " +
		"FROM organization " +
		"WHERE id = ? and status = ?"
)
