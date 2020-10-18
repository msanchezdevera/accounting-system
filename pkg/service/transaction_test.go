package service

import (
	"accounting/api/transaction"
	"accounting/pkg/log"
	"accounting/pkg/model"
	"accounting/pkg/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

var creditTransactionAPI = &transaction.CreateTransaction{
	Amount: 100,
	Type:   "credit",
}

var creditTransaction = &model.Transaction{
	TransactionType: model.Credit,
	Amount:          100,
}

func TestTransactionService(t *testing.T) {
	transactionStorage := storage.NewTransaction()

	accountService := NewAccountService()

	service := NewTransactionService(transactionStorage, accountService, log.NewConfigless())

	response, err := service.Create(creditTransactionAPI)
	creditTransaction.ID = response.ID
	creditTransaction.EffectiveDate = response.EffectiveDate

	assert.NoError(t, err)
	assert.Equal(t, response, creditTransaction)

	response = service.Get(creditTransaction.ID)
	assert.Equal(t, response, creditTransaction)

	all := service.GetAll()
	assert.Equal(t, all, []*model.Transaction{creditTransaction})
}
