package main

import (
	"context"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func (node *Node) buildP2pServer() {
	node.P2pServer = &http.Server{
		Handler: websocket.Handler(func(ws *websocket.Conn) {
			node.WebSocketConnections = append(node.WebSocketConnections, ws)
			node.handleP2pConnection(ws)
		}),
		Addr: ":6001",
	}
}

func (node *Node) runP2pServer() {
	go func() {
		fmt.Println("WebSocket server for P2P communication is listening on ws://localhost:6001")
		if err := node.P2pServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
}

func (node *Node) shutdownP2pServer () {
	fmt.Println("Shutting down WebSocket server.")
	node.P2pServer.Shutdown(context.Background())
}

func (node *Node) handleP2pConnection(conn *websocket.Conn) {
	// TODO
}

