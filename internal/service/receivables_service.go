package service

import (
	"fmt"
	"recorderData/entity/receivables"
)

type ReceivableService struct {
	Repository receivables.ReceivableRepository
}

func NewReceivableService(repository receivables.ReceivableRepository) *ReceivableService {
	return &ReceivableService{Repository: repository}
}

func (r *ReceivableService) findById(id string) (receivables.Receivable, error) {
	receivableInstance, err := r.Repository.GetReceivable(id)
	if err != nil {
		return receivables.Receivable{}, err
	}
	return receivableInstance, nil
}

func (r *ReceivableService) Create(name string, receivableType string, amount float64, date string, status string,
	description string) (receivables.Receivable, error) {

	receivableInstance := receivables.NewReceivable(name, receivableType, amount, date, status, description)
	fmt.Println("Instance: ", receivableInstance)
	receivableSaved, err := r.Repository.SaveReceivable(*receivableInstance)
	if err != nil {
		return receivables.Receivable{}, err
	}
	return receivableSaved, nil
}
