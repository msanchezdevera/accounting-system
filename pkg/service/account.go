package service

import (
	"accounting/pkg/model"
	"sync"
)

type AccountService interface {
	Balance() float64
	LockBalance() float64
	UpdateBalance(newAmount float64)
	UnlockBalance()
}

type accountService struct {
	account model.Account
	rwMutex *sync.RWMutex
}

func NewAccountService() AccountService {
	return &accountService{
		rwMutex: &sync.RWMutex{},
		account: model.Account{},
	}
}

func (acc *accountService) Balance() float64 {
	acc.rwMutex.RLock()
	defer acc.rwMutex.RUnlock()
	return acc.account.Balance
}

func (acc *accountService) LockBalance() float64 {
	acc.rwMutex.Lock()
	return acc.account.Balance
}

func (acc *accountService) UpdateBalance(newAmount float64) {
	acc.account.Balance = newAmount
}

func (acc *accountService) UnlockBalance() {
	acc.rwMutex.Unlock()
}
