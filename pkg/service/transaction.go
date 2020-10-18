package service

import (
	"accounting/api/transaction"
	"accounting/pkg/errors"
	"accounting/pkg/log"
	"accounting/pkg/model"
	"accounting/pkg/storage"
	"github.com/google/uuid"
	"time"
)

type TransactionService interface {
	Create(create *transaction.CreateTransaction) (*model.Transaction, errors.Error)
	Get(transactionId string) *model.Transaction
	GetAll() []*model.Transaction
}

func NewTransactionService(transactionStorage storage.Transaction, account *model.Account, log log.Logger) TransactionService {
	return &transactionService{
		transactionStorage: transactionStorage,
		log:                log,
		account:            account,
	}
}

type transactionService struct {
	transactionStorage storage.Transaction
	log                log.Logger
	account            *model.Account
}

func (ts *transactionService) Create(transactionCreate *transaction.CreateTransaction) (*model.Transaction, errors.Error) {
	ts.log.Infof("Creating transaction: %v", transactionCreate)

	var newUuid uuid.UUID
	var err error

	if newUuid, err = uuid.NewRandom(); err != nil {
		return nil, errors.New(err.Error())
	}

	transactionModel := &model.Transaction{
		ID:              newUuid.String(),
		TransactionType: model.TransactionType(transactionCreate.Type),
		Amount:          transactionCreate.Amount,
		EffectiveDate:   time.Now(),
	}

	balance := ts.account.LockBalance()
	defer ts.account.UnlockBalance()

	if transactionModel.TransactionType == model.Credit {
		balance += transactionModel.Amount
	} else if transactionModel.TransactionType == model.Debit {
		balance -= transactionModel.Amount
	}

	ts.account.UpdateBalance(balance)

	ts.transactionStorage.Create(transactionModel)

	return transactionModel, nil
}

func (ts *transactionService) Get(transactionId string) *model.Transaction {
	return ts.transactionStorage.Get(transactionId)
}

func (ts *transactionService) GetAll() []*model.Transaction {
	return ts.transactionStorage.GetAll()
}
