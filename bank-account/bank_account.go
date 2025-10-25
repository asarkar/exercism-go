package account

import (
	"sync"
)

// Define the Account type here.
type Account struct {
	balance int64
	open    bool
	mu      sync.RWMutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, true, sync.RWMutex{}}
}

func (account *Account) Close() (int64, bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.open {
		return 0, false
	}
	account.open = false
	b := account.balance
	account.balance = 0
	return b, true
}

func (account *Account) Balance() (int64, bool) {
	account.mu.RLock()
	defer account.mu.RUnlock()
	return account.balance, account.open
}

func (account *Account) Deposit(amount int64) (int64, bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.open || account.balance+amount < 0 {
		return account.balance, false
	}
	account.balance += amount
	return account.balance, true
}
