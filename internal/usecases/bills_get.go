package usecases

import (
	"database/sql"
	"log"
	"recorderData/entity/bills"
	"recorderData/internal/repository"
	"recorderData/internal/service"
)

func GetBill(billDraft bills.Bill) (bills.Bill, error) {

	dbBills, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqliteBills := repository.NewRepositoryBillsSqlite(dbBills)

	billsService := service.NewBillsService(repositorySqliteBills)
	billGot, err := billsService.FindById(billDraft.Id)
	if err != nil {
		return bills.Bill{}, nil
	}
	log.Default().Println("Bill got: ", billGot)
	return billGot, nil
}
