package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"recorderData/entity/receivables"
	"recorderData/internal/dto"
	"recorderData/internal/usecases"
)

func ReceivableGetHandler(w http.ResponseWriter, r *http.Request) {

	var receivableDraft receivables.Receivable

	json.NewDecoder(r.Body).Decode(&receivableDraft)

	receivableGot, err := usecases.GetReceivable(receivableDraft)
	if err != nil {
		log.Fatal("get use case: " + err.Error())
	}

	receivableDTO, err := dto.NewReceivable(receivableGot)
	if err != nil {
		log.Fatal("get receivable dto case: " + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.MarshalIndent(receivableDTO, "", "  ")
	if err != nil {
		log.Fatal("marshal receivable dto: " + err.Error())
	}

	w.Write(responseJSON)
}

func ReceivablePostHandler(w http.ResponseWriter, r *http.Request) {

	var receivableDraft receivables.Receivable

	json.NewDecoder(r.Body).Decode(&receivableDraft)
	receivableSaved, err := usecases.CreateReceivable(receivableDraft)
	if err != nil {
		log.Fatal("create use case: " + err.Error())
	}

	fmt.Fprintf(w, "receivable saved: %v", receivableSaved)
}
