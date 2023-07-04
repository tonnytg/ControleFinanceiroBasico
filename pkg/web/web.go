package web

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/stocks", StocksHandler)
	http.HandleFunc("/bills", BillsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func StocksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		StockGetHandler(w, r)
	}

	if r.Method == "POST" {
		StockPostHandler(w, r)
	}
}

func BillsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		BillGetHandler(w, r)
	}

	if r.Method == "POST" {
		BillPostHandler(w, r)
	}
}
