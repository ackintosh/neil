package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"strconv"
)

func (node *Node) buildP2pServer() {
	node.P2pServer = &http.Server{
		Handler: websocket.Handler(func(ws *websocket.Conn) {
			node.WebSocketConnections = append(node.WebSocketConnections, ws)
			node.handleP2pConnection(ws)
		}),
		Addr: ":" + strconv.Itoa(*p2pPort),
	}
}

func (node *Node) runP2pServer() {
	go func() {
		fmt.Println("WebSocket server for P2P communication is listening on ws://localhost" + node.P2pServer.Addr)
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
	for {
		var message []byte
		if err := websocket.Message.Receive(conn, &message); err != nil {
			if err == io.EOF {
				node.disconnect(conn)
				break
			}
			fmt.Println(err)
			continue
		}

		fmt.Println(string(message))
	}
}

func (node *Node) broadcast(message *Message) {
	marshaledMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Failed to marshal a message: ", err)
	}

	for _, conn := range node.WebSocketConnections {
		if err := websocket.Message.Send(conn, marshaledMessage); err != nil {
			fmt.Println(err)
		}
	}
}

func (node *Node) disconnect(conn *websocket.Conn) {
	defer conn.Close()
	fmt.Println("disconnect: ", conn.RemoteAddr().String())
	node.deleteWsConnection(conn)
}

func (node *Node) deleteWsConnection(conn *websocket.Conn) {
	conns := []*websocket.Conn{}
	for _, c := range node.WebSocketConnections {
		if c != conn {
			conns = append(conns, c)
		}
	}

	node.WebSocketConnections = conns
}
