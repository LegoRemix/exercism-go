// Package account simulates a bank account
package account

import "sync"

// Account is a datatype representing a bank account, synchronized for concurrent access
type Account struct {
	sync.RWMutex       // all accounts are implicitly RWMutexes
	balance      int64 // the current amount of money in the account
	invalid      bool  // is this account invalid?
}

// Open attempts to create an account with a specified initialDeposit, fails if initialDeposit < 0
func Open(initialDeposit int64) *Account {
	// if we have less than 0 money, fail to open an account
	if initialDeposit < int64(0) {
		return nil
	}
	account := new(Account)
	account.balance = initialDeposit

	return account
}

// Balance attempts to get the account balance for a provided account (reading operation)
func (acc *Account) Balance() (balance int64, ok bool) {
	acc.RLock()
	defer acc.RUnlock()

	//if the account is invalid, just return now
	if acc.invalid {
		return
	}

	ok = true
	balance = acc.balance

	return
}

// Close attempts to close out an account and render it invalid (writing operation)
func (acc *Account) Close() (payout int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	// if the account is already invalidated, just bail
	if acc.invalid {
		return
	}

	//here we invalidate the account
	ok = true
	acc.invalid = true
	payout = acc.balance

	return
}

// Deposit attempts to update the amount of money in an account (writing operation)
func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	//make sure the account is valid AND that we aren't overdrafting
	if acc.invalid || amount+acc.balance < 0 {
		return
	}

	ok = true
	newBalance = amount + acc.balance
	acc.balance = newBalance

	return
}
