package p2p

import (
	"bufio"
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
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

func handleStream(s network.Stream) {
	fmt.Printf("New stream received from %s\n", s.Conn().RemotePeer().ShortString())
	reader := bufio.NewReader(s)
	msg, _ := reader.ReadString('\n')
	fmt.Printf("Message received: %s", msg)
}

// CreateNode は、新しいlibp2pノードを作成し、それをネットワークに追加します。
func (nm *NodeManager) CreateNode(pid protocol.ID) host.Host {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	node.SetStreamHandler(pid, handleStream)

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

// ListNodes は、ネットワーク内のすべてのノードのIDとアドレスを表示します。
func (nm *NodeManager) ListNodes() {
	nm.lock.Lock()
	defer nm.lock.Unlock()
	for i, node := range nm.nodes {
		fmt.Printf("%d: node ID: %s\n", i+1, node.ID().ShortString())
	}
}
