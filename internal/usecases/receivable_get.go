package usecases

import (
	"database/sql"
	"log"
	"recorderData/entity/receivables"
	"recorderData/internal/repository"
	"recorderData/internal/service"
)

func GetReceivable(receivableDraft receivables.Receivable) (receivables.Receivable, error) {

	dbReceivable, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqliteReceivable := repository.NewRepositoryReceivablesSqlite(dbReceivable)

	receivableService := service.NewReceivableService(repositorySqliteReceivable)
	receivableGot, err := receivableService.FindById(receivableDraft.Id)
	if err != nil {
		return receivables.Receivable{}, nil
	}
	log.Default().Println("Receivable got: ", receivableGot)
	return receivableGot, nil
}
