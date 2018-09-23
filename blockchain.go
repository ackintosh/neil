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

func (bc *Blockchain) NewBlock(prevBlockHash []byte, transactions []*Transaction) *Block {
	block := &Block{
		len(bc.blocks),
		time.Now().Unix(),
		prevBlockHash,
		[]byte{},
		transactions,
		0,
	}
	block.SetHash()
	return block
}

func (bc *Blockchain) CreateGenesisBlock() {
	bc.blocks = append(bc.blocks, bc.NewBlock([]byte{}, []*Transaction{}))
}

func (bc *Blockchain) AddBlock() {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := bc.NewBlock(prevBlock.Hash, bc.transactions)
	bc.transactions = []*Transaction{}
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) AddTransaction(tx *Transaction) {
	bc.transactions = append(bc.transactions, tx)
}
