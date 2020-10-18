package model

import "sync"

type Account struct {
	balance float64
	rwMutex *sync.RWMutex
}

func NewAccount() *Account {
	return &Account{
		rwMutex: &sync.RWMutex{},
		balance: 0,
	}
}

func (acc *Account) Balance() float64 {
	acc.rwMutex.RLock()
	defer acc.rwMutex.RUnlock()
	return acc.balance
}

func (acc *Account) LockBalance() float64 {
	acc.rwMutex.Lock()
	return acc.balance
}

func (acc *Account) UpdateBalance(newAmount float64) {
	acc.balance = newAmount
}

func (acc *Account) UnlockBalance() {
	acc.rwMutex.Unlock()
}
