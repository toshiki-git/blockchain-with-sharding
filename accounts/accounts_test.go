package accounts

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	account := NewAccount("testAddress", 1000)

	if account.Address != "testAddress" {
		t.Errorf("Expected address to be 'testAddress', got %s", account.Address)
	}

	if account.Balance != 1000 {
		t.Errorf("Expected balance to be 1000, got %d", account.Balance)
	}
}
