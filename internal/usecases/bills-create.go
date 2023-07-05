package usecases

import (
	"database/sql"
	"recorderData/entity/bills"
	"recorderData/internal/repository"
	"recorderData/internal/service"
)

func CreateBill(billDraft bills.Bill) (bills.Bill, error) {

	dbBills, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqliteBills := repository.NewRepositoryBillsSqlite(dbBills)

	bill, err := bills.NewBill(billDraft)
	if err != nil {
		return bills.Bill{}, nil
	}

	billsService := service.NewBillsService(repositorySqliteBills)
	stockSaved, err := billsService.Create(bill)
	if err != nil {
		return bills.Bill{}, nil
	}
	return stockSaved, nil
}
