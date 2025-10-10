package transaction

import (
	"database/sql"
)

type TransactionManagerImpl struct {
	db *sql.DB
}

func (tm *TransactionManagerImpl) WithTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := tm.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func NewTransactionManager(db *sql.DB) TransactionManager {
	return &TransactionManagerImpl{
		db: db,
	}
}
