package service

import (
	"fmt"
	"recorderData/entity/bills"
)

type BillsService struct {
	Repository bills.BillRepository
}

func NewBillsService(repository bills.BillRepository) *BillsService {
	return &BillsService{Repository: repository}
}

func (b *BillsService) FindById(id string) (bills.Bill, error) {
	billInstance, err := b.Repository.GetBill(id)
	if err != nil {
		return bills.Bill{}, err
	}
	return billInstance, nil
}

func (b *BillsService) Create(currentBill bills.Bill) (bills.Bill, error) {

	billInstance, err := bills.NewBill(currentBill)
	if err != nil {
		return bills.Bill{}, err
	}
	fmt.Println("Instance: ", billInstance)
	billSaved, err := b.Repository.SaveBill(billInstance)
	if err != nil {
		return bills.Bill{}, err
	}
	return billSaved, nil
}
