package main

import (
	"database/sql"
	"fmt"
	"log"
	"recorderData/entity/bills"
	"recorderData/entity/receivables"
	"recorderData/entity/stocks"
	"recorderData/internal/repository"
	"recorderData/internal/service"
)

func main() {
	dbReceivable := repository.ReceivableMemoryDb{Receivable: []receivables.Receivable{}}
	repositoryMemory := repository.NewRepositoryReceivableMemory(dbReceivable)

	receivableService := service.NewReceivableService(repositoryMemory)
	r, _ := receivableService.Create("teste", "teste", 10.0, "teste", "teste", "teste")
	fmt.Println(r)

	dbReceivable2, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqlite := repository.NewRepositorySqlite(dbReceivable2)

	receivableService2 := service.NewReceivableService(repositorySqlite)
	rr, err := receivableService2.Create("teste1", "teste1", 11.0, "teste1", "teste1", "teste1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rr)

	dbBill1 := repository.BillsMemoryDb{Bills: []bills.Bill{}}
	repositoryMemoryBill := repository.NewRepositoryBillsMemory(dbBill1)

	billService := service.NewBillsService(repositoryMemoryBill)
	myBillDraft := bills.Bill{
		Title:          "teste1",
		Receiver:       "teste1",
		Description:    "teste1",
		Value:          100,
		Tax:            0,
		DateIssued:     "teste1",
		DateExpiration: "teste1",
		Observation:    "teste1",
	}
	myBill, _ := bills.NewBill(myBillDraft)
	b, _ := billService.Create(myBill)
	fmt.Println(b)

	dbBill2, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqliteBill := repository.NewRepositoryBillsSqlite(dbBill2)

	billService2 := service.NewBillsService(repositorySqliteBill)
	myBill2, err := bills.NewBill(myBillDraft)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := billService2.Create(myBill2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b2)

	dbStock, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqliteStock := repository.NewRepositoryStocksSqlite(dbStock)

	stockDraft := stocks.Stock{
		Name:        "teste",
		Description: "teste",
		Amount:      10.0,
		Quantity:    10.0,
		Tax:         10.0,
		DateMade:    "teste",
		Action:      "teste",
	}

	stock, _ := stocks.NewStock(stockDraft)

	stockService := service.NewStockssService(repositorySqliteStock)
	stockSaved, err := stockService.Create(stock)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stockSaved)

}
