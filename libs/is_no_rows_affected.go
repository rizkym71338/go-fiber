package libs

func IsNoRowsAffected(rowsAffected int64) bool {
	return rowsAffected == 0
}
