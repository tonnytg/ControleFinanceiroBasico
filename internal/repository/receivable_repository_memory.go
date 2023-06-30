package repository

import (
	"errors"
	"recorderData/entity/receivables"
)

type ReceivableMemoryDb struct {
	Receivable []receivables.Receivable
}

type ReceivableRepositoryMemory struct {
	db ReceivableMemoryDb
}

func NewRepositoryReceivableMemory(db ReceivableMemoryDb) *ReceivableRepositoryMemory {
	return &ReceivableRepositoryMemory{db: db}
}

func (r *ReceivableRepositoryMemory) GetReceivable(id string) (receivables.Receivable, error) {
	for _, receivable := range r.db.Receivable {
		if receivable.Id == id {
			return receivable, nil
		}
	}
	return receivables.Receivable{}, errors.New("receivable not found with this Id")
}

func (r *ReceivableRepositoryMemory) SaveReceivable(receivable receivables.Receivable) (receivables.Receivable, error) {
	r.db.Receivable = append(r.db.Receivable, receivable)
	return receivable, nil
}
