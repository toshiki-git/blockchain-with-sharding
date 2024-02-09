package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/toshiki-git/blockchain-with-sharding/p2p"
)

func main() {
	ctx := context.Background()
	nodes := make([]host.Host, 4)

	// 4つのノードを作成
	for i := range nodes {
		node := p2p.CreateNode()
		nodes[i] = node
		fmt.Printf("ノード%d: %s\n", i, node.ID().ShortString())
	}

	// ノードを相互に接続
	for i, node := range nodes {
		for j, otherNode := range nodes {
			if i == j {
				continue
			}
			addrInfo := peer.AddrInfo{
				ID:    otherNode.ID(),
				Addrs: otherNode.Addrs(),
			}
			if err := node.Connect(ctx, addrInfo); err != nil {
				panic(err)
			}
		}
	}

	// メッセージの送信ロジック
	sendRandomMessage := func(nodes []host.Host) {
		senderIndex := rand.Intn(len(nodes))
		receiverIndex := rand.Intn(len(nodes))
		// 自分自身には送信しないようにする
		for senderIndex == receiverIndex {
			receiverIndex = rand.Intn(len(nodes))
		}

		sender := nodes[senderIndex]
		receiver := nodes[receiverIndex]

		fmt.Printf("メッセージ送信: %s -> %s\n", sender.ID().ShortString(), receiver.ID().ShortString())

		// 実際にはここでメッセージを送信します
		// この例では、実際のメッセージ送信メカニズムは実装されていません
	}

	// 一定間隔でメッセージを送信
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			sendRandomMessage(nodes)
		}
	}()

	// 無限ループでプログラムが終了しないようにする
	select {}
}
