package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
)

type Node struct {
	Blockchain *Blockchain
	ApiServer *http.Server
}

func NewNode() *Node {
	blockchain := NewBlockchain()
	blockchain.AddTransaction(NewTransaction([]byte("Bob"), []byte("Ivan"), 1))
	blockchain.AddTransaction(NewTransaction([]byte("Bob"), []byte("Ivan"), 2))

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

func (node *Node) runMining() {
	fmt.Println("Mining blocks...")
	go func () {
		for {
			node.ProofOfWork()
		}
	}()
}

func (node *Node) ProofOfWork() {
	block := node.Blockchain.createBlock()
	var nonce int64 = 0
	var hash [32]byte
	for {
		headers := bytes.Join(
			[][]byte{
				block.PrevBlockHash[:],// [32]byte -> []byte
				[]byte(strconv.FormatInt(block.Timestamp, 10)),
				[]byte(strconv.FormatInt(nonce, 10)),
			},
			[]byte{},
		)
		hash = sha256.Sum256(headers)
		if bytes.Equal(hash[:2], []byte("00")) {
			break
		}
		nonce++
	}

	block.Nonce = nonce
	block.Hash = hash[:]
	node.Blockchain.blocks = append(node.Blockchain.blocks, block)
}
