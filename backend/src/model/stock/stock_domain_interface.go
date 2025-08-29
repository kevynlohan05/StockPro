package model

type StockDomainInterface interface {
	GetID() string
	GetProductID() int
	GetUserID() int
	GetStockType() string
	GetQuantity() int
	GetReason() string

	SetID(id string)
}

func NewStockDomain(productID int, userID int, stockType string, quantity int, reason string) StockDomainInterface {
	return &StockDomain{
		productID: productID,
		userID:    userID,
		stockType: stockType,
		quantity:  quantity,
		reason:    reason,
	}
}
