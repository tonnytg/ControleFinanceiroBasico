package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"recorderData/entity/bills"
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

	billDTO, err := dto.NewBill(billGot)
	if err != nil {
		log.Fatal("create bill dto case: " + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.MarshalIndent(billDTO, "", "  ")
	if err != nil {
		log.Fatal("marshal stock dto: " + err.Error())
	}

	w.Write(responseJSON)
}

func BillPostHandler(w http.ResponseWriter, r *http.Request) {
	var billDraft bills.Bill
	json.NewDecoder(r.Body).Decode(&billDraft)

	billSaved, err := usecases.CreateBill(billDraft)
	if err != nil {
		log.Fatal("create use case: " + err.Error())
	}
	fmt.Fprintf(w, "Stock saved: %v", billSaved)
}
