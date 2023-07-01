package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"recorderData/entity/stocks"
	"recorderData/internal/usecases"
)

func Start() {

	http.HandleFunc("/stocks", CreateStockHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func CreateStockHandler(w http.ResponseWriter, r *http.Request) {

	var stockDraft stocks.Stock
	json.NewDecoder(r.Body).Decode(&stockDraft)

	stockSaved, err := usecases.CreateStock(stockDraft)
	if err != nil {
		log.Fatal("create use case: " + err.Error())
	}
	fmt.Fprintf(w, "Stock saved: %v", stockSaved)

}
