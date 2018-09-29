package main

import "time"

type Chain struct {
	blocks []*Block
	transactions []*Transaction
}

func NewChain() *Chain {
	bc := &Chain{[]*Block{}, []*Transaction{}}
	bc.CreateGenesisBlock()
	return bc
}

func (bc *Chain) NewBlock(prevBlockHash string, transactions []*Transaction) *Block {
	block := &Block{
		len(bc.blocks),
		time.Now().Unix(),
		prevBlockHash,
		"",
		transactions,
		0,
	}
	return block
}

func (bc *Chain) CreateGenesisBlock() {
	genesisBlock := bc.NewBlock("0", []*Transaction{})
	genesisBlock.Hash = "000001ec215dbddf866b9ac3219d8d33875ff0c19ca88ab84cf83b2e3397b489"
	bc.blocks = append(bc.blocks, genesisBlock)
}

func (bc *Chain) createBlock() *Block {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := bc.NewBlock(prevBlock.Hash, bc.transactions)
	bc.transactions = []*Transaction{}

	return newBlock
}

func (bc *Chain) getLatestBlock() *Block {
	return bc.blocks[len(bc.blocks) - 1]
}

func (bc *Chain) AddTransaction(tx *Transaction) {
	bc.transactions = append(bc.transactions, tx)
}
