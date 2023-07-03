package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"recorderData/entity/stocks"
	"recorderData/internal/dto"
	"recorderData/internal/usecases"
)

func StockGetHandler(w http.ResponseWriter, r *http.Request) {
	var stockDraft stocks.Stock
	json.NewDecoder(r.Body).Decode(&stockDraft)

	stockGot, err := usecases.GetStock(stockDraft)
	if err != nil {
		log.Fatal("create use case: " + err.Error())
	}

	stockDTO, err := dto.NewStock(stockGot)
	if err != nil {
		log.Fatal("create stock dto case: " + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.MarshalIndent(stockDTO, "", "  ")
	if err != nil {
		log.Fatal("marshal stock dto: " + err.Error())
	}

	w.Write(responseJSON)
}

func StockPostHandler(w http.ResponseWriter, r *http.Request) {
	var stockDraft stocks.Stock
	json.NewDecoder(r.Body).Decode(&stockDraft)

	stockSaved, err := usecases.CreateStock(stockDraft)
	if err != nil {
		log.Fatal("create use case: " + err.Error())
	}
	fmt.Fprintf(w, "Stock saved: %v", stockSaved)
}
