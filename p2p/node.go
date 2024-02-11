package p2p

import (
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
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

// CreateNode は、新しいlibp2pノードを作成し、それをネットワークに追加します。
func (nm *NodeManager) CreateNode() host.Host {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	node.SetStreamHandler("/echo/1.0.0", func(s network.Stream) {
		fmt.Println("接続が確立されました")
	})
	nm.lock.Lock()
	defer nm.lock.Unlock()
	nm.nodes = append(nm.nodes, node)
	return node
}

// ConnectNodes は、２つのノード間で接続を確立します。
func (nm *NodeManager) ConnectNodes(node1, node2 host.Host) error {
	node2Info := peer.AddrInfo{
		ID:    node2.ID(),
		Addrs: node2.Addrs(),
	}
	return node1.Connect(context.Background(), node2Info)
}

// ListNodes は、ネットワーク内のすべてのノードのIDとアドレスを表示します。
func (nm *NodeManager) ListNodes() {
	nm.lock.Lock()
	defer nm.lock.Unlock()
	for _, node := range nm.nodes {
		addrs, err := peer.AddrInfoToP2pAddrs(&peer.AddrInfo{
			ID:    node.ID(),
			Addrs: node.Addrs(),
		})
		if err != nil {
			fmt.Println("アドレスの取得に失敗しました:", err)
			continue
		}
		fmt.Printf("node ID: %s, address: %s\n\n", node.ID().ShortString(), addrs)
	}
}
