package main

type Block struct {
	Index int
	Timestamp int64
	PrevBlockHash string
	Hash string
	Transactions []*Transaction
	Nonce int64
}
