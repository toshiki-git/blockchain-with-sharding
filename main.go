package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

// 　メッセージを受信したときの処理
/* func handleStream(s network.Stream) {
	fmt.Printf("New stream received from %s\n", s.Conn().RemotePeer().ShortString())
	reader := bufio.NewReader(s)
	msg, _ := reader.ReadString('\n')
	fmt.Printf("Message received: %s", msg)
}
*/

func handleStream(s network.Stream) {
	fmt.Printf("New stream received from %s\n", s.Conn().RemotePeer().ShortString())
	reader := bufio.NewReader(s)

	// メッセージを1行読み込む（JSONデータを想定）
	msg, err := reader.ReadString('\n')
	if err != nil {
		panic(err) // 実際には適切にエラーハンドリングを行う
	}

	// 受け取ったメッセージをBlock型にデコード
	var block Block
	if err := json.Unmarshal([]byte(msg), &block); err != nil {
		panic(err) // 実際には適切にエラーハンドリングを行う
	}

	fmt.Printf("Block received: %+v\n", block)
}

type Block struct {
	Index     int
	Timestamp string
	CreateBy  string
}

func main() {
	ctx := context.Background()

	// ノード1の作成
	node1, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	defer node1.Close()

	// ノード2の作成
	node2, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	if err != nil {
		panic(err)
	}
	node2.SetStreamHandler("/myapp/1.0.0", handleStream)
	defer node2.Close()

	// ノード2の作成
	node3, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	if err != nil {
		panic(err)
	}
	defer node3.Close()

	// ノード2のアドレス情報を出力
	for _, addr := range node2.Addrs() {
		fmt.Printf("Node2 is listening on %s/p2p/%s\n", addr, node2.ID().ShortString())
	}

	fmt.Printf("Node2 ID111111: %s\n", node2.Addrs())

	/* peerInfo1 := peer.AddrInfo{
		ID:    node1.ID(),
		Addrs: node1.Addrs(),
	} */

	peerInfo2 := peer.AddrInfo{
		ID:    node2.ID(),
		Addrs: node2.Addrs(),
	}

	/* peerInfo3 := peer.AddrInfo{
		ID:    node3.ID(),
		Addrs: node3.Addrs(),
	} */

	// ノード1からノード2へ接続
	if err := node1.Connect(ctx, peerInfo2); err != nil {
		panic(err)
	}

	if err := node3.Connect(ctx, peerInfo2); err != nil {
		panic(err)
	}

	// node1に接続されているすべてのピアのIDを取得し、表示する
	peers := node2.Network().Peers()
	for _, peerID := range peers {
		fmt.Printf("Node2 is connected to: %s\n", peerID.ShortString())
	}

	// ノード1からノード2へメッセージを送信
	s, err := node1.NewStream(ctx, node2.ID(), "/myapp/1.0.0")
	if err != nil {
		panic(err)
	}
	block := Block{Index: 2, Timestamp: "2021-01-01", CreateBy: "node3"}
	blockBytes, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}
	_, err = s.Write(blockBytes)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(s) // JSONデータの後に改行を追加
	if err := s.Close(); err != nil {
		panic(err)
	}

	// ノード3からノード2へメッセージを送信
	s, err = node3.NewStream(ctx, node2.ID(), "/myapp/1.0.0")
	if err != nil {
		panic(err)
	}

	block = Block{Index: 1, Timestamp: "2021-01-01", CreateBy: "node3"}
	blockBytes, err = json.Marshal(block)
	if err != nil {
		panic(err)
	}
	_, err = s.Write(blockBytes)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(s) // JSONデータの後に改行を追加
	if err := s.Close(); err != nil {
		panic(err)
	}

	/* fmt.Fprintln(s, block)
	if err := s.Close(); err != nil {
		panic(err)
	} */

	// プログラムが終了しないように待つ
	select {}
}
