package stocks

import "github.com/google/uuid"

type StockRepository interface {
	GetStock(id string) (Stock, error)
	SaveStock(stock Stock) (Stock, error)
}

type Stock struct {
	Id          string
	Name        string
	Description string
	Amount      float64
	Quantity    float64
	Tax         float64
	DateMade    string
	Action      string
}

func NewStock(stock Stock) (Stock, error) {

	stock = Stock{
		Id:          uuid.NewString(),
		Name:        stock.Name,
		Description: stock.Description,
		Amount:      stock.Amount,
		Quantity:    stock.Quantity,
		Tax:         stock.Tax,
		DateMade:    stock.DateMade,
		Action:      stock.Action,
	}
	return stock, nil
}
