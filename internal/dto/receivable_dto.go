package dto

import (
	"recorderData/entity/receivables"
)

type ReceivableDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
}

func NewReceivable(receivable receivables.Receivable) (ReceivableDTO, error) {

	var receivableDTO = ReceivableDTO{
		ID:          receivable.Id,
		Name:        receivable.Name,
		Type:        receivable.Type,
		Amount:      receivable.Amount,
		Date:        receivable.Date,
		Status:      receivable.Status,
		Description: receivable.Description,
	}

	return receivableDTO, nil
}
