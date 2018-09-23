package main

type Block struct {
	Index int
	Timestamp int64
	PrevBlockHash [32]byte
	Hash [32]byte
	Transactions []*Transaction
	Nonce int64
}
