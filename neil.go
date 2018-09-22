package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	node := NewNode()

	for _, block := range node.Blockchain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	node.RunApiServer()

	fmt.Println("Neil is running...")
	fmt.Println("Press Ctrl+C to stop. :)")

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := <-signalCh
	fmt.Println("\nsignal: ", s)
	node.ShutdownApiServer()

	fmt.Println("\nNeil has completely shutdown. Thanks for giving it a try!")
}
