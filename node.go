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
	blockchain.AddTransaction(NewTransaction([]byte("Bob"), []byte("Ivan"), 1))
	blockchain.AddTransaction(NewTransaction([]byte("Bob"), []byte("Ivan"), 2))
	blockchain.AddBlock("Sample block")

	node := &Node{blockchain, nil}
	node.buildApiServer()

	return node
}

func (node *Node) RunApiServer() {
	go func() {
		fmt.Println("REST API server is listening on http://localhost:3001")
		if err := node.ApiServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
}

func (node *Node) ShutdownApiServer() {
	fmt.Println("Shutting down REST API server.")
	node.ApiServer.Shutdown(context.Background())
}
