package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Transaction struct {
	Sender []byte
	Recipient []byte
	Amount int64
	Hash string
}

func NewTransaction(sender []byte, recipient []byte, amount int64) *Transaction {
	tx := &Transaction{sender, recipient, amount, ""}
	tx.setHash()
	return tx
}

func (tx *Transaction) setHash() {
	headers := bytes.Join(
		[][]byte{
			tx.Sender,
			tx.Recipient,
			[]byte(strconv.FormatInt(tx.Amount, 10)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(headers)
	tx.Hash = hex.EncodeToString(hash[:])
}
