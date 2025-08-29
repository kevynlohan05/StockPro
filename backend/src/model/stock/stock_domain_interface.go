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

type StockQuantityDomainInterface interface {
	GetQuantityStock() int
	GetProductIDStock() int
}

func NewStockDomain(productID int, userID int, stockType string, quantity int, reason string) StockDomainInterface {
	return &movementStockDomain{
		productID: productID,
		userID:    userID,
		stockType: stockType,
		quantity:  quantity,
		reason:    reason,
	}
}

func NewStockEntity(productID int, userID int, stockType string, quantity int, reason string, created_at int) StockDomainInterface {
	return &movementStockDomain{
		productID:  productID,
		userID:     userID,
		stockType:  stockType,
		quantity:   quantity,
		reason:     reason,
		created_at: created_at,
	}
}

func NewStockQuantityEntity(productID, quantity int) StockQuantityDomainInterface {
	return &stockDomain{
		productID: productID,
		quantity:  quantity,
	}
}
