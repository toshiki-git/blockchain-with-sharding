package consensus

import (
	"math/rand"
	"time"

	"github.com/toshiki-git/blockchain-with-sharding/accounts"
	"github.com/toshiki-git/blockchain-with-sharding/blockchain"
)

type Consensus struct {
	Chain          []*blockchain.Block
	AccountManager *accounts.AccountManager
}

// AddBlock はブロックをチェーンに追加する
func (c *Consensus) AddBlock(block *blockchain.Block) {
	c.Chain = append(c.Chain, block)
}

// ValidateBlock はブロックの有効性を検証する
func (c *Consensus) ValidateBlock(block *blockchain.Block) bool {
	// ここにはブロックの検証ロジックを実装します
	// 例: 前のブロックのハッシュが正しいか、トランザクションが有効かなど
	return true
}

// CreateBlock は新しいブロックを作成する
func (c *Consensus) CreateBlock(transactions []*blockchain.Transaction) *blockchain.Block {
	// ランダムにアカウントを選択
	rand.Seed(time.Now().UnixNano())
	accountAddresses := make([]string, 0, len(c.AccountManager.Accounts))
	for address := range c.AccountManager.Accounts {
		accountAddresses = append(accountAddresses, address)
	}
	randomAccount := c.AccountManager.Accounts[accountAddresses[rand.Intn(len(accountAddresses))]]

	// ランダムに選ばれたアカウントを使用してブロックを作成
	var lastHash string
	if len(c.Chain) > 0 {
		lastHash = c.Chain[len(c.Chain)-1].Header.PreviousHash
	}

	block := blockchain.NewBlock(int64(len(c.Chain)), lastHash, transactions, randomAccount.Address)
	if c.ValidateBlock(block) {
		c.AddBlock(block)
		return block
	}
	return nil
}
