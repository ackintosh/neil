package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGTERM)
	<-signalCh
}
