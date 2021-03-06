package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	apiPort     = flag.Int("apiPort", 3001, "REST API server port")
	p2pPort     = flag.Int("p2pPort", 6001, "WebSocket server port for P2P communication")
	passiveMode = flag.Bool("passive", false, "The node this flag passed doesn't mine new block.")
)

func main() {
	flag.Parse()
	node := NewNode()
	node.runApiServer()
	node.runP2pServer()
	if *passiveMode {
		fmt.Println("This node is running as 'passive' mode, mining is disabled.")
	} else {
		node.runMining()
	}

	fmt.Println("Neil is running...")
	fmt.Println("Press Ctrl+C to stop. :)")

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := <-signalCh
	fmt.Println("\nsignal: ", s)
	node.shutdownApiServer()
	node.shutdownP2pServer()

	fmt.Println("\nNeil has completely shutdown. Thanks for giving it a try!")
}
