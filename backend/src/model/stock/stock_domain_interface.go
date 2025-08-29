package model

type StockDomainInterface interface {
	GetID() string
	GetProductID() int
	GetUserID() int
	GetStockType() string
	GetQuantity() int
	GetReason() string
	GetCreatedAt() int

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

func NewStockEntity(productID int, userID int, stockType string, quantity int, reason string, created_at int) StockDomainInterface {
	return &StockDomain{
		productID:  productID,
		userID:     userID,
		stockType:  stockType,
		quantity:   quantity,
		reason:     reason,
		created_at: created_at,
	}
}
