package accounts

import (
	"fmt"
	"strconv"
)

type Account struct {
	Address string
	Balance int
}

type AccountManager struct {
	Accounts map[string]*Account
	Count    int
}

func InitializeWithDummyAccounts() *AccountManager {
	manager := NewAccountManager()

	for i := 1; i <= 10; i++ {
		address := "Address_" + strconv.Itoa(i)
		balance := 1000
		account := NewAccount(address, balance)
		manager.AddAccount(account)
	}

	return manager
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

func NewAccountManager() *AccountManager {
	return &AccountManager{
		Accounts: make(map[string]*Account),
		Count:    0,
	}
}

func (m *AccountManager) AddAccount(account *Account) {
	m.Accounts[account.Address] = account
	m.Count++
}

func (m *AccountManager) DisplayAccounts() {
	for _, account := range m.Accounts {
		account.Display()
	}
	fmt.Printf("Total Accounts: %d\n", m.Count)
}
