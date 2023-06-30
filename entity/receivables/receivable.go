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

func NewReceivable(Name string, Type string, Amount float64, Date string, Status string, Description string) *Receivable {
	r := Receivable{
		Id:          uuid.NewString(),
		Name:        Name,
		Type:        Type,
		Amount:      Amount,
		Date:        Date,
		Status:      Status,
		Description: Description,
	}
	return &r
}
