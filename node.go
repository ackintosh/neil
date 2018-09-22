package main

type Node struct {
	Blockchain *Blockchain
}

func NewNode() *Node {
	blockchain := NewBlockchain()
	blockchain.AddBlock("Send 1 BTC to Ivan")
	blockchain.AddBlock("Send 2 more BTC to Ivan")

	return &Node{blockchain}
}
