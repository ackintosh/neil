package main

type Transaction struct {
	sender []byte
	recipient []byte
	amount int64
}

func NewTransaction(sender []byte, recipient []byte, amount int64) *Transaction {
	return &Transaction{sender, recipient, amount}
}
