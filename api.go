package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
)

func (node *Node) buildApiServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/chain", node.chainHandler)
	mux.HandleFunc("/transactions/new", node.newTransactionHandler)
	mux.HandleFunc("/peers", node.peersHandler)

	node.ApiServer = &http.Server{
		Handler: mux,
		Addr: ":" + strconv.Itoa(*apiPort),
	}
}

func (node *Node) runApiServer() {
	go func() {
		fmt.Println("REST API server is listening on http://localhost" + node.ApiServer.Addr)
		if err := node.ApiServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
}

func (node *Node) shutdownApiServer() {
	fmt.Println("Shutting down REST API server.")
	node.ApiServer.Shutdown(context.Background())
}

func (node *Node) chainHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(node.Chain.blocks)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (node *Node) newTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Sender []byte `json:"Sender"`
		Recipient []byte `json:"Recipient"`
		Amount int64 `json:"Amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	node.Chain.AddTransaction(NewTransaction(params.Sender, params.Recipient, params.Amount))
	w.WriteHeader(http.StatusCreated)
}

func (node *Node) peersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		node.postPeersHandler(w, r)
		return
	} else if r.Method == http.MethodGet {
		node.getPeersHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (node *Node) postPeersHandler(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Address string `json:"Address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ws, err := websocket.Dial(params.Address, "", "http://127.0.0.1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connected to the peer(triggered by the REST API `/peers`): ", ws.RemoteAddr().String())
	node.WebSocketConnections = append(node.WebSocketConnections, ws)

	node.Peers = append(node.Peers, params.Address)
	w.WriteHeader(http.StatusCreated)
}

func (node *Node) getPeersHandler(w http.ResponseWriter, r *http.Request) {
	var peersArray []map[string]interface{}
	for _, c := range node.WebSocketConnections {
		peersArray = append(peersArray, map[string]interface{}{"Address": c.RemoteAddr().String()})
	}
	peersJson, err := json.Marshal(peersArray)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(peersJson)
}
