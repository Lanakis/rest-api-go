package utils

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		NewAppError(errRollback, "Cant rollback", "")
	} else {
		errCommit := tx.Commit()
		NewAppError(errCommit, "Cant rollback", "")
	}
}
