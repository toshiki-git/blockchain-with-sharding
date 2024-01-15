package accounts

import (
	"fmt"
)

type Account struct {
	Address string
	Balance int
}

func NewAccount(address string, initialBalance int) *Account {
	return &Account{
		Address: address,
		Balance: initialBalance,
	}
}

func (a *Account) Display() {
	fmt.Printf("Address: %s\nBalance: %d\n", a.Address, a.Balance)
}
