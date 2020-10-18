package model

import (
	"accounting/pkg/errors"
	"time"
)

type TransactionType string

type Transaction struct {
	ID              string
	TransactionType TransactionType
	Amount          float64
	EffectiveDate   time.Time
}

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

func (tt TransactionType) IsValid() errors.Error {
	switch tt {
	case Credit, Debit:
		return nil
	}

	return errors.UserError.New("Invalid transaction type")
}
