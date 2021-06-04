package tx

import "database/sql"

type Transaction struct {
	DBTransaction *sql.Tx
}

func (t Transaction) Commit() error {
	return t.DBTransaction.Commit()
}

func (t Transaction) Rollback() error {
	return t.DBTransaction.Rollback()
}

type TransactionFactory interface {
	NewTransaction() (Transaction, error)
}

var _ TransactionFactory = (*SafeTransactionFactory)(nil)

type SafeTransactionFactory struct {
	db *sql.DB
}

func (s SafeTransactionFactory) NewTransaction() (Transaction, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{DBTransaction: tx}, nil
}

func NewSafeTransactionFactory(db *sql.DB) SafeTransactionFactory {
	return SafeTransactionFactory{db: db}
}
