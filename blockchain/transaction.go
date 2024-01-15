package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Transaction struct {
	ID               string
	SenderAddress    string
	RecipientAddress string
	Amount           int64
}

func NewTransaction(sender, recipient string, amount int64) *Transaction {
	return &Transaction{
		ID:               generateTransactionID(sender, recipient, amount),
		SenderAddress:    sender,
		RecipientAddress: recipient,
		Amount:           amount,
	}
}

func generateTransactionID(sender, recipient string, amount int64) string {
	timestamp := time.Now().UnixNano()
	input := fmt.Sprintf("%s:%s:%d:%d", sender, recipient, timestamp, amount)

	hash := sha256.Sum256([]byte(input))

	return fmt.Sprintf("%x", hash)
}
