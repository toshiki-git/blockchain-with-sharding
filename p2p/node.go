package p2p

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/toshiki-git/blockchain-with-sharding/blockchain"
)

// NodeManager は、ネットワーク内のノードを管理します。
type NodeManager struct {
	nodes []host.Host
	lock  sync.Mutex
}

// NewNodeManager は、新しいNodeManagerインスタンスを作成します。
func NewNodeManager() *NodeManager {
	return &NodeManager{}
}

func handleMessageStream(s network.Stream) {
	fmt.Printf("New Messagestream received from %s\n", s.Conn().RemotePeer().ShortString())
	reader := bufio.NewReader(s)
	msg, _ := reader.ReadString('\n')
	fmt.Printf("Message received: %s", msg)
}

func handleBlockStream(s network.Stream) {
	fmt.Printf("New Blockstream received from %s\n", s.Conn().RemotePeer().ShortString())

	// io.ReadAllを使用して、ストリームからデータを全て読み込む
	data, err := io.ReadAll(s)
	if err != nil {
		fmt.Printf("Failed to read data from stream: %s\n", err)
		return
	}

	// 読み込んだデータをブロック構造体にアンマーシャル
	var block blockchain.Block
	if err := json.Unmarshal(data, &block); err != nil {
		fmt.Printf("Error unmarshalling block: %s\n", err)
		return
	}

	fmt.Printf("Block received: %+v\n", block)
}

// CreateNode は、新しいlibp2pノードを作成し、それをネットワークに追加します。
func (nm *NodeManager) CreateNode() host.Host {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	node.SetStreamHandler("/message/1.0.0", handleMessageStream)
	node.SetStreamHandler("/block/1.0.0", handleBlockStream)

	nm.lock.Lock()
	defer nm.lock.Unlock()
	nm.nodes = append(nm.nodes, node)
	return node
}

// ConnectNodes は、２つのノード間で接続を確立します。
func (nm *NodeManager) ConnectNode(node1, node2 host.Host) error {
	node2Info := peer.AddrInfo{
		ID:    node2.ID(),
		Addrs: node2.Addrs(),
	}
	return node1.Connect(context.Background(), node2Info)
}

func (nm *NodeManager) SendMessage(from, to host.Host, pid protocol.ID) {
	s, err := from.NewStream(context.Background(), to.ID(), pid)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(s, "This is message from %s to %s", from.ID().ShortString(), to.ID().ShortString())
	if err := s.Close(); err != nil {
		panic(err)
	}
}

func (nm *NodeManager) SendBlock(from, to host.Host, pid protocol.ID) {
	s, err := from.NewStream(context.Background(), to.ID(), pid)
	if err != nil {
		panic(err)
	}

	header := blockchain.BlockHeader{Number: 1, PreviousHash: "0x11111", Timestamp: 20240212}
	transactions := []*blockchain.Transaction{
		{
			ID:               "tx1",
			SenderAddress:    "Alice",
			RecipientAddress: "Bob",
			Amount:           100,
		},
		{
			ID:               "tx2",
			SenderAddress:    "Charlie",
			RecipientAddress: "Dave",
			Amount:           50,
		},
	}

	block := blockchain.Block{Header: header, Transactions: transactions, Creator: "CreatorAddress"}
	blockBytes, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}
	_, err = s.Write(blockBytes)
	if err != nil {
		panic(err)
	}
	if err := s.Close(); err != nil {
		panic(err)
	}
}

// ListNodes は、ネットワーク内のすべてのノードのIDとアドレスを表示します。
func (nm *NodeManager) ListNodes() {
	nm.lock.Lock()
	defer nm.lock.Unlock()
	for i, node := range nm.nodes {
		fmt.Printf("%d: node ID: %s\n", i+1, node.ID().ShortString())
	}
}
