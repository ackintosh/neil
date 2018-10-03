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
	NewChainCh chan NewChainEvent
}

func NewNode() *Node {
	chain := NewChain()
	chain.AddTransaction(NewTransaction([]byte("Bob"), []byte("Ivan"), 1))
	chain.AddTransaction(NewTransaction([]byte("Bob"), []byte("Ivan"), 2))

	node := &Node{
		chain,
		nil,
		nil,
		[]*websocket.Conn{},
		make(chan NewChainEvent, 1),
	}
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

MINING:
	for {
		select {
		case event := <-node.NewChainCh:
			fmt.Println("Received a new block. we may restart current mining process. ", event)
		default:
			hash = node.calculateBlockHash(block, nonce)
			if node.isMeetCriteria(hash) {
				break MINING
			}
			nonce++
		}
	}

	block.Nonce = nonce
	block.Hash = hash
	node.Chain.blocks = append(node.Chain.blocks, block)
	fmt.Println("Added new block: " + block.Hash)

	message, err := newLatestBlockMessage(block)
	if err != nil {
		fmt.Println("Failed to build a message: ", err)
		return
	}

	node.broadcast(message)
}

func (node *Node) calculateBlockHash(block *Block, nonce int64) string {
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
	return hex.EncodeToString(hash32Bytes[:])
}

func (node *Node) isMeetCriteria(blockHash string) bool {
	return blockHash[:6] == "000000"
}
