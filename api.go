package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (node *Node) buildApiServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/chain", node.chainHandler)
	mux.HandleFunc("/transactions/new", node.newTransactionHandler)
	mux.HandleFunc("/nodes", node.nodesHandler)

	node.ApiServer = &http.Server{
		Handler: mux,
		Addr: ":3001",
	}
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

func (node *Node) nodesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		node.postNodesHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (node *Node) postNodesHandler(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Address string `json:"Address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	node.Nodes = append(node.Nodes, params.Address)
	w.WriteHeader(http.StatusCreated)
}
