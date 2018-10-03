package main

import "encoding/json"

type Message struct {
	Type MessageType
	Data string
}

type MessageType int

const (
	MessageTypeLatestBlock MessageType = iota
	MessageTypeNewTransaction MessageType = iota
)

func newLatestBlockMessage(latestBlock *Block) (*Message, error) {
	bytes, err := json.Marshal(latestBlock)
	if err != nil {
		return nil, err
	}

	return &Message{
		Type: MessageTypeLatestBlock,
		Data: string(bytes),
	}, nil
}

func newNewTransactionMessage(transaction *Transaction) (*Message, error) {
	bytes, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}

	return &Message{
		Type: MessageTypeNewTransaction,
		Data: string(bytes),
	}, nil
}
