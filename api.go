package main

import "net/http"

func (node *Node) buildApiServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/blocks", node.blocksHandler)

	node.ApiServer = &http.Server{
		Handler: mux,
		Addr: ":3001",
	}
}

func (node *Node) blocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte{})
}
