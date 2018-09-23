package main

import "time"

type Blockchain struct {
	blocks []*Block
	transactions []*Transaction
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{[]*Block{}, []*Transaction{}}
	bc.CreateGenesisBlock()
	return bc
}

func (bc *Blockchain) NewBlock(data string, prevBlockHash []byte, transactions []*Transaction) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, transactions}
	block.SetHash()
	return block
}

func (bc *Blockchain) CreateGenesisBlock() {
	bc.blocks = append(bc.blocks, bc.NewBlock("Genesis Block", []byte{}, []*Transaction{}))
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := bc.NewBlock(data, prevBlock.Hash, bc.transactions)
	bc.transactions = []*Transaction{}
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) AddTransaction(tx *Transaction) {
	bc.transactions = append(bc.transactions, tx)
}
