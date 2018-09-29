package main

import "encoding/json"

type Message struct {
	Type MessageType
	Data string
}

type MessageType int

const (
	MessageTypeLatestBlock MessageType = iota
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
