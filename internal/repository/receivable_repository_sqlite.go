package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"recorderData/entity/receivables"
)

type ReceivableRepositorySqlite struct {
	db *sql.DB
}

func NewRepositoryReceivablesSqlite(db *sql.DB) *ReceivableRepositorySqlite {
	return &ReceivableRepositorySqlite{db: db}
}

func (r *ReceivableRepositorySqlite) GetReceivable(id string) (receivables.Receivable, error) {
	var receivableInstance receivables.Receivable

	stmt, err := r.db.Prepare("SELECT id, name FROM receivable WHERE id = ?")
	if err != nil {
		return receivableInstance, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&receivableInstance.Id,
		&receivableInstance.Name,
	)
	if err != nil {
		log.Default().Println("Receivable got: ", receivableInstance)
		return receivableInstance, err
	}

	log.Default().Println("Receivable got: ", receivableInstance)
	return receivableInstance, nil
}

func (r *ReceivableRepositorySqlite) SaveReceivable(receivable receivables.Receivable) (receivables.Receivable, error) {
	log.Default().Println("Saving receivable in sqlite repository")
	stmt, err := r.db.Prepare("insert into receivable (id, name) values (?, ?)")
	if err != nil {
		return receivable, err
	}

	fmt.Println("Receivable received: ", receivable)
	_, err = stmt.Exec(receivable.Id, receivable.Name)
	if err != nil {
		return receivable, err
	}
	err = stmt.Close()
	if err != nil {
		return receivable, err
	}
	return receivable, nil
}
