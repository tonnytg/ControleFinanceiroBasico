package service

import (
	"recorderData/entity/receivables"
)

type ReceivableService struct {
	Repository receivables.ReceivableRepository
}

func NewReceivableService(repository receivables.ReceivableRepository) *ReceivableService {
	return &ReceivableService{Repository: repository}
}

func (r *ReceivableService) FindById(id string) (receivables.Receivable, error) {

	receivableInstance, err := r.Repository.GetReceivable(id)
	if err != nil {
		return receivables.Receivable{}, err
	}

	return receivableInstance, nil
}

func (r *ReceivableService) Create(receivable receivables.Receivable) (receivables.Receivable, error) {

	receivableInstance, err := receivables.NewReceivable(receivable)
	if err != nil {
		return receivables.Receivable{}, err
	}

	receivableSaved, err := r.Repository.SaveReceivable(receivableInstance)
	if err != nil {
		return receivables.Receivable{}, err
	}

	return receivableSaved, nil
}
