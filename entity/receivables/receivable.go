package receivables

import (
	"github.com/google/uuid"
)

type ReceivableRepository interface {
	GetReceivable(id string) (Receivable, error)
	SaveReceivable(receivable Receivable) (Receivable, error)
}

type Receivable struct {
	Id          string
	Name        string
	Type        string
	Amount      float64
	Date        string
	Status      string
	Description string
}

func NewReceivable(receivable Receivable) (Receivable, error) {
	r := Receivable{
		Id:          uuid.NewString(),
		Name:        receivable.Name,
		Type:        receivable.Type,
		Amount:      receivable.Amount,
		Date:        receivable.Date,
		Status:      receivable.Status,
		Description: receivable.Description,
	}
	return r, nil
}
