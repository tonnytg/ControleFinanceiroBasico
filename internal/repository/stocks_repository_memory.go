package repository

import (
	"errors"
	"recorderData/entity/stocks"
)

type StocksMemoryDb struct {
	Stocks []stocks.Stock
}

type StocksRepositoryMemory struct {
	db StocksMemoryDb
}

func NewRepositoryStocksMemory(db StocksMemoryDb) *StocksRepositoryMemory {
	return &StocksRepositoryMemory{db: db}
}

func (s *StocksRepositoryMemory) GetStock(id string) (stocks.Stock, error) {
	for _, stock := range s.db.Stocks {
		if stock.Id == id {
			return stock, nil
		}
	}
	return stocks.Stock{}, errors.New("stock not found with this Id")
}

func (s *StocksRepositoryMemory) SaveStock(bill stocks.Stock) (stocks.Stock, error) {
	s.db.Stocks = append(s.db.Stocks, bill)
	return bill, nil
}
