package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"recorderData/entity/bills"
	"recorderData/entity/stocks"
	"recorderData/internal/dto"
	"recorderData/internal/usecases"
)

func BillGetHandler(w http.ResponseWriter, r *http.Request) {
	var bill bills.Bill
	json.NewDecoder(r.Body).Decode(&bill)

	billGot, err := usecases.GetBill(bill)
	if err != nil {
		log.Fatal("create use case: " + err.Error())
	}

	stockDTO, err := dto.NewStock(billGot)
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
