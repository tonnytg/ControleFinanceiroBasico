package usecases

import (
	"database/sql"
	"recorderData/entity/stocks"
	"recorderData/internal/repository"
	"recorderData/internal/service"
)

func CreateStock(stockDraft stocks.Stock) (stocks.Stock, error) {

	dbStock, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqliteStock := repository.NewRepositoryStocksSqlite(dbStock)

	stock, err := stocks.NewStock(stockDraft)
	if err != nil {
		return stocks.Stock{}, nil
	}

	stockService := service.NewStockssService(repositorySqliteStock)
	stockSaved, err := stockService.Create(stock)
	if err != nil {
		return stocks.Stock{}, nil
	}
	return stockSaved, nil
}
