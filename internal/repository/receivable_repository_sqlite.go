package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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

	stmt, err := r.db.Prepare("SELECT * FROM receivable WHERE id = ?")
	if err != nil {
		return receivableInstance, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&receivableInstance.Id,
		&receivableInstance.Name,
		&receivableInstance.Type,
		&receivableInstance.Amount,
		&receivableInstance.Date,
		&receivableInstance.Status,
		&receivableInstance.Description,
	)
	if err != nil {
		return receivableInstance, err
	}

	return receivableInstance, nil
}

func (r *ReceivableRepositorySqlite) SaveReceivable(receivable receivables.Receivable) (receivables.Receivable, error) {

	stmt, err := r.db.Prepare("insert into receivable (id, name) values (?, ?)")
	if err != nil {
		return receivable, err
	}

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
