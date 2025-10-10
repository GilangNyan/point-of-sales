package transaction

import (
	"database/sql"
)

type TransactionManager interface {
	WithTransaction(fn func(tx *sql.Tx) error) error
}
