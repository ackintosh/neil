package main

type Block struct {
	Index int
	Timestamp int64
	PrevBlockHash []byte
	Hash []byte
	Transactions []*Transaction
	Nonce int64
}
