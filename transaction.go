package main

type Transaction struct {
	Sender []byte
	Recipient []byte
	Amount int64
	Hash []byte
}

func NewTransaction(sender []byte, recipient []byte, amount int64) *Transaction {
	return &Transaction{sender, recipient, amount, []byte{}}
}
