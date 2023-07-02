package usecases

import (
	"database/sql"
	"log"
	"recorderData/entity/stocks"
	"recorderData/internal/repository"
	"recorderData/internal/service"
)

func GetStock(stockDraft stocks.Stock) (stocks.Stock, error) {

	dbStock, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqliteStock := repository.NewRepositoryStocksSqlite(dbStock)

	stockService := service.NewStockssService(repositorySqliteStock)
	stockGot, err := stockService.FindById(stockDraft.Id)
	if err != nil {
		return stocks.Stock{}, nil
	}
	log.Default().Println("Stock got: ", stockGot)
	return stockGot, nil
}
