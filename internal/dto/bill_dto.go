package dto

import "recorderData/entity/bills"

type BillDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Quantity    float64 `json:"quantity"`
	Tax         float64 `json:"tax"`
	DateMade    string  `json:"dateMade"`
	Action      string  `json:"action"`
}

func NewBill(bill bills.Bill) (BillDTO, error) {

	var billDTO = BillDTO{
		ID:          bill.Id,
		Name:        bill.Title,
		Description: bill.Description,
		Amount:      bill.Value,
		Quantity:    bill.Tax,
		Tax:         bill.Tax,
		DateMade:    bill.DateIssued,
		Action:      bill.Observation,
	}
	return billDTO, nil
}
