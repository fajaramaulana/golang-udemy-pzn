package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	// recover digunakan untuk menangkap panic
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		// rollback akan mengembalikan perubahan data ke semula
		PanicIfError(errorRollback)
	} else {
		errorCommit := tx.Commit()
		// commit akan menyimpan perubahan data
		PanicIfError(errorCommit)
	}
}
