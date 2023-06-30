package repository

import (
	"errors"
	"recorderData/entity/bills"
)

type BillsMemoryDb struct {
	Bills []bills.Bill
}

type BillsRepositoryMemory struct {
	db BillsMemoryDb
}

func NewRepositoryBillsMemory(db BillsMemoryDb) *BillsRepositoryMemory {
	return &BillsRepositoryMemory{db: db}
}

func (r *BillsRepositoryMemory) GetBill(id string) (bills.Bill, error) {
	for _, bill := range r.db.Bills {
		if bill.Id == id {
			return bill, nil
		}
	}
	return bills.Bill{}, errors.New("bill not found with this Id")
}

func (r *BillsRepositoryMemory) SaveBill(bill bills.Bill) (bills.Bill, error) {
	r.db.Bills = append(r.db.Bills, bill)
	return bill, nil
}
