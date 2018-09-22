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

	apiServer := NewApiServer()
	go func() {
		fmt.Println("API server is running on port 3001.")
		if err := apiServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Neil is running...")
	fmt.Println("Press Ctrl+C to stop. :)")

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGTERM)
	<-signalCh
}
