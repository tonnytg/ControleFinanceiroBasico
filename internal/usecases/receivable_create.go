package usecases

import (
	"database/sql"
	"recorderData/entity/receivables"
	"recorderData/internal/repository"
	"recorderData/internal/service"
)

func CreateReceivable(receivableDraft receivables.Receivable) (receivables.Receivable, error) {

	dbReceivables, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqliteReceivable := repository.NewRepositoryReceivablesSqlite(dbReceivables)

	receivable, err := receivables.NewReceivable(receivableDraft)
	if err != nil {
		return receivables.Receivable{}, nil
	}

	receivableService := service.NewReceivableService(repositorySqliteReceivable)
	stockSaved, err := receivableService.Create(receivable)
	if err != nil {
		return receivables.Receivable{}, nil
	}
	return stockSaved, nil
}
