package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (node *Node) buildApiServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/blocks", node.blocksHandler)

	node.ApiServer = &http.Server{
		Handler: mux,
		Addr: ":3001",
	}
}

func (node *Node) blocksHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(node.Blockchain.blocks)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
