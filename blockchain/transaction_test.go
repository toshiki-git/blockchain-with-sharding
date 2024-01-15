package blockchain

import (
	"testing"
)

func TestNewTransaction(t *testing.T) {
	transaction := NewTransaction("testSender", "testRecipient", 100)

	if transaction.SenderAddress != "testSender" {
		t.Errorf("Expected sender address to be 'testSender', got %s", transaction.SenderAddress)
	}

	if transaction.RecipientAddress != "testRecipient" {
		t.Errorf("Expected recipient address to be 'testRecipient', got %s", transaction.RecipientAddress)
	}

	if transaction.Amount != 100 {
		t.Errorf("Expected amount to be 100, got %d", transaction.Amount)
	}
}

func TestGenerateTransactionID(t *testing.T) {
	transactionID := generateTransactionID("testSender", "testRecipient", 100)

	if transactionID == "" {
		t.Errorf("Expected transaction ID to be non-empty, got %s", transactionID)
	}
}
