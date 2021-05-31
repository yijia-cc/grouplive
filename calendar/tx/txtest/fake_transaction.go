package txtest

import "github.com/yijia-cc/grouplive/calendar/tx"

var _ tx.TransactionFactory = (*FakeTransactionFactory)(nil)

type FakeTransactionFactory struct {
}

func (f FakeTransactionFactory) NewTransaction() (tx.Transaction, error) {
	return tx.Transaction{}, nil
}

func NewFakeTransactionFactory() FakeTransactionFactory {
	return FakeTransactionFactory{}
}

