package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type BlockHeader struct {
	Number       int64
	PreviousHash string
	Timestamp    int64
}

type Block struct {
	Header       BlockHeader
	Transactions []*Transaction
	Creator      string // 生成者のアドレスを追加
}

func NewBlock(number int64, previousHash string, transactions []*Transaction, creator string) *Block {
	block := &Block{
		Header: BlockHeader{
			Number:       number,
			PreviousHash: previousHash,
			Timestamp:    time.Now().Unix(),
		},
		Transactions: transactions,
		Creator:      creator,
	}

	blockHash := calculateBlockHash(block)
	block.Header.PreviousHash = blockHash

	return block
}

func calculateBlockHash(block *Block) string {
	record := fmt.Sprintf("%d%s%d", block.Header.Number, block.Header.PreviousHash, block.Header.Timestamp)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
