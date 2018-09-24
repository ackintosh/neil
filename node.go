package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
)

type Node struct {
	Chain                *Chain
	ApiServer            *http.Server
	P2pServer            *http.Server
	WebSocketConnections []*websocket.Conn
	Peers                []string
}

func NewNode() *Node {
	chain := NewChain()
	chain.AddTransaction(NewTransaction([]byte("Bob"), []byte("Ivan"), 1))
	chain.AddTransaction(NewTransaction([]byte("Bob"), []byte("Ivan"), 2))

	node := &Node{chain, nil, nil, []*websocket.Conn{},[]string{}}
	node.buildApiServer()
	node.buildP2pServer()

	return node
}

func (node *Node) runMining() {
	fmt.Println("Mining blocks...")
	go func () {
		for {
			node.proofOfWork()
		}
	}()
}

func (node *Node) proofOfWork() {
	block := node.Chain.createBlock()
	var nonce int64 = 0
	var hash string
	for {
		var txHashes [][]byte
		for _, tx := range block.Transactions {
			txHashes = append(txHashes, []byte(tx.Hash))
		}
		headers := bytes.Join(
			[][]byte{
				[]byte(block.PrevBlockHash),
				bytes.Join(txHashes, []byte{}),
				[]byte(strconv.FormatInt(block.Timestamp, 10)),
				[]byte(strconv.FormatInt(nonce, 10)),
			},
			[]byte{},
		)
		hash32Bytes := sha256.Sum256(headers)
		hash = hex.EncodeToString(hash32Bytes[:])
		if hash[:5] == "00000" {
			break
		}
		nonce++
	}

	block.Nonce = nonce
	block.Hash = hash
	node.Chain.blocks = append(node.Chain.blocks, block)
	fmt.Print("Added new block: ")
	fmt.Println(block.Hash)
}
