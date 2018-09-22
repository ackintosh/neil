package main

import (
	"context"
	"fmt"
	"net/http"
)

type Node struct {
	Blockchain *Blockchain
	ApiServer *http.Server
}

func NewNode() *Node {
	blockchain := NewBlockchain()
	blockchain.AddBlock("Send 1 BTC to Ivan")
	blockchain.AddBlock("Send 2 more BTC to Ivan")

	return &Node{blockchain, NewApiServer()}
}

func (node *Node) RunApiServer() {
	go func() {
		fmt.Println("API server is running on port 3001.")
		if err := node.ApiServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
}

func (node *Node) ShutdownApiServer() {
	fmt.Println("Shutting down API server.")
	node.ApiServer.Shutdown(context.Background())
}
