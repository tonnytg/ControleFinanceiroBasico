package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"recorderData/entity/bills"
)

type BillsRepositorySqlite struct {
	db *sql.DB
}

func NewRepositoryBillsSqlite(db *sql.DB) *BillsRepositorySqlite {
	return &BillsRepositorySqlite{db: db}
}

func (b *BillsRepositorySqlite) GetBill(id string) (bills.Bill, error) {

	var billInstance bills.Bill
	stmt, err := b.db.Prepare("select * from bills where id = ?")
	if err != nil {
		return billInstance, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&billInstance.Id,
		&billInstance.Title,
		&billInstance.Receiver,
		&billInstance.Description,
		&billInstance.Value,
		&billInstance.Tax,
		&billInstance.DateIssued,
		&billInstance.DateExpiration,
		&billInstance.Observation,
	)
	if err != nil {
		return billInstance, errors.New("bill not found with this Id")
	}
	return billInstance, nil
}

func (b *BillsRepositorySqlite) SaveBill(bill bills.Bill) (bills.Bill, error) {
	log.Default().Println("Saving bills in sqlite repository")
	fmt.Println("bill: ", bill)
	stmt, err := b.db.Prepare("insert into bills (id, title, receiver, description, amount, tax, dateissued, dataexperiation, observation) values (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return bill, err
	}

	fmt.Println("Bill received: ", bill)
	_, err = stmt.Exec(bill.Id, bill.Title, bill.Receiver, bill.Description, bill.Value, bill.Tax, bill.DateIssued, bill.DateExpiration, bill.Observation)
	if err != nil {
		return bill, err
	}
	err = stmt.Close()
	if err != nil {
		return bill, err
	}
	return bill, nil
}
