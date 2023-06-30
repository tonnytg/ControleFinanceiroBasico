package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"recorderData/entity/stocks"
)

type StocksRepositorySqlite struct {
	db *sql.DB
}

func NewRepositoryStocksSqlite(db *sql.DB) *StocksRepositorySqlite {
	return &StocksRepositorySqlite{db: db}
}

func (s *StocksRepositorySqlite) GetStock(id string) (stocks.Stock, error) {

	var stockInstance stocks.Stock
	stmt, err := s.db.Prepare("select * from stocks where id = ?")
	if err != nil {
		return stockInstance, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&stockInstance.Id, &stockInstance.Name)
	if err != nil {
		return stockInstance, errors.New("stock not found with this Id")
	}
	return stockInstance, nil
}

func (s *StocksRepositorySqlite) SaveStock(stock stocks.Stock) (stocks.Stock, error) {
	log.Default().Println("Saving stock in sqlite repository")
	fmt.Println("bill: ", stock)
	stmt, err := s.db.Prepare("insert into stocks (id, name, description, amount, quantity, tax, datemade, action) values (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return stock, err
	}

	fmt.Println("Stock received: ", stock)
	_, err = stmt.Exec(stock.Id, stock.Name, stock.Description, stock.Amount, stock.Quantity, stock.Tax, stock.DateMade, stock.Action)
	if err != nil {
		return stock, err
	}
	err = stmt.Close()
	if err != nil {
		return stock, err
	}
	return stock, nil
}
