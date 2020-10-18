package storage

import (
	"accounting/pkg/model"
	"sync"
)

type Transaction interface {
	Create(transaction *model.Transaction)
	Get(id string) *model.Transaction
	GetAll() []*model.Transaction
}

type transaction struct {
	transactions *sync.Map
}

func NewTransaction() Transaction {
	return &transaction{
		transactions: &sync.Map{},
	}
}

func (t *transaction) Create(transaction *model.Transaction) {
	t.transactions.Store(transaction.ID, transaction)
}

func (t *transaction) Get(id string) *model.Transaction {
	if result, found := t.transactions.Load(id); found {
		return result.(*model.Transaction)
	}

	return nil
}

func (t *transaction) GetAll() []*model.Transaction {
	var transactions []*model.Transaction
	t.transactions.Range(func(key, value interface{}) bool {
		transactions = append(transactions, value.(*model.Transaction))
		return true
	})

	return transactions
}
