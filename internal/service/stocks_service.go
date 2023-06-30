package service

import (
	"fmt"
	"recorderData/entity/stocks"
)

type StocksService struct {
	Repository stocks.StockRepository
}

func NewStockssService(repository stocks.StockRepository) *StocksService {
	return &StocksService{Repository: repository}
}

func (s *StocksService) findById(id string) (stocks.Stock, error) {
	stockInstance, err := s.Repository.GetStock(id)
	if err != nil {
		return stocks.Stock{}, err
	}
	return stockInstance, nil
}

func (s *StocksService) Create(stock stocks.Stock) (stocks.Stock, error) {

	stockInstance, err := stocks.NewStock(stock)
	if err != nil {
		return stocks.Stock{}, err
	}
	fmt.Println("Instance: ", stockInstance)
	stockSaved, err := s.Repository.SaveStock(stockInstance)
	if err != nil {
		return stocks.Stock{}, err
	}
	return stockSaved, nil
}
