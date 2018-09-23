package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Index int
	Timestamp int64
	PrevBlockHash []byte
	Hash []byte
	Transactions []*Transaction
	Nonce int64
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
