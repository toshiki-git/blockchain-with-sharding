package main

import (
	"github.com/toshiki-git/blockchain-with-sharding/p2p"
)

func main() {
	manager := p2p.NewNodeManager()
	node1 := manager.CreateNode()
	node2 := manager.CreateNode()

	manager.ConnectNode(node1, node2)

	manager.ListNodes()

	//manager.SendMessage(node1, node2, "/message/1.0.0")
	manager.SendBlock(node1, node2, "/block/1.0.0")

	select {}
}
