package bills

import "github.com/google/uuid"

type BillRepository interface {
	GetBill(id string) (Bill, error)
	SaveBill(bill Bill) (Bill, error)
}

type Bill struct {
	Id             string
	Title          string
	Receiver       string
	Description    string
	Value          float64
	Tax            float64
	DateIssued     string
	DateExpiration string
	Observation    string
}

func NewBill(currentBill Bill) (Bill, error) {

	var bill = Bill{
		Id:             uuid.NewString(),
		Title:          currentBill.Title,
		Receiver:       currentBill.Receiver,
		Description:    currentBill.Description,
		Value:          currentBill.Value,
		Tax:            currentBill.Tax,
		DateIssued:     currentBill.DateIssued,
		DateExpiration: currentBill.DateExpiration,
		Observation:    currentBill.Observation,
	}
	return bill, nil
}
