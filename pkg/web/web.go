package web

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/stocks", StockHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func StockHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		StockGetHandler(w, r)
	}

	if r.Method == "POST" {
		StockPostHandler(w, r)
	}
}
