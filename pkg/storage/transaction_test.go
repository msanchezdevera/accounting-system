package storage

import (
	"accounting/pkg/model"
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

var creditTransaction = &model.Transaction{
	ID:              "test-id",
	TransactionType: model.Credit,
	Amount:          100,
	EffectiveDate:   time.Time{},
}

func TestTransactionStorage(t *testing.T) {
	transactionStorage := NewTransaction()
	transactionStorage.Create(creditTransaction)

	response := transactionStorage.Get(creditTransaction.ID)
	assert.Equal(t, response, creditTransaction)

	allTransactions := transactionStorage.GetAll()
	assert.Equal(t, allTransactions, []*model.Transaction{creditTransaction})
}
