package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	node := NewNode()
	node.RunApiServer()
	node.runMining()

	fmt.Println("Neil is running...")
	fmt.Println("Press Ctrl+C to stop. :)")

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := <-signalCh
	fmt.Println("\nsignal: ", s)
	node.ShutdownApiServer()

	fmt.Println("\nNeil has completely shutdown. Thanks for giving it a try!")
}
