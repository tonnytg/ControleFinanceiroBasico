package dto

import (
	"recorderData/entity/stocks"
)

type StockDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Quantity    float64 `json:"quantity"`
	Tax         float64 `json:"tax"`
	DateMade    string  `json:"dateMade"`
	Action      string  `json:"action"`
}

func NewStock(stock stocks.Stock) (StockDTO, error) {
	var stockDTO StockDTO
	stockDTO.ID = stock.Id
	stockDTO.Name = stock.Name
	stockDTO.Description = stock.Description
	stockDTO.Amount = stock.Amount
	stockDTO.Quantity = stock.Quantity
	stockDTO.Tax = stock.Tax
	stockDTO.DateMade = stock.DateMade
	stockDTO.Action = stock.Action

	return stockDTO, nil
}
