package main

import (
	"github.com/toshiki-git/blockchain-with-sharding/p2p"
)

func main() {
	manager := p2p.NewNodeManager()
	node1 := manager.CreateNode("/echo/1.0.0")
	node2 := manager.CreateNode("/echo/1.0.0")

	manager.ConnectNode(node1, node2)

	manager.ListNodes()

	manager.SendMessage(node1, node2, "/echo/1.0.0")

	select {}
}
