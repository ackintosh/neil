package main

type Blockchain struct {
	blocks []*Block
	transactions []*Transaction
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, bc.transactions)
	bc.transactions = []*Transaction{}
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) AddTransaction(tx *Transaction) {
	bc.transactions = append(bc.transactions, tx)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}, []*Transaction{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}, []*Transaction{}}
}