package main

import "net/http"

func NewApiServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/blocks", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte{})
	})

	return &http.Server{
		Handler: mux,
		Addr: ":3001",
	}
}
