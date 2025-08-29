package model

type movementStockDomain struct {
	iD         string
	productID  int
	userID     int
	stockType  string
	quantity   int
	reason     string
	created_at int
}

type stockDomain struct {
	productID int
	quantity  int
}

func (sd *stockDomain) GetProductIDStock() int {
	return sd.productID
}

func (sd *stockDomain) GetQuantityStock() int {
	return sd.quantity
}

func (sd *movementStockDomain) GetID() string {
	return sd.iD
}

func (sd *movementStockDomain) GetProductID() int {
	return sd.productID
}

func (sd *movementStockDomain) GetUserID() int {
	return sd.userID
}

func (sd *movementStockDomain) GetStockType() string {
	return sd.stockType
}

func (sd *movementStockDomain) GetQuantity() int {
	return sd.quantity
}

func (sd *movementStockDomain) GetReason() string {
	return sd.reason
}

func (sd *movementStockDomain) GetCreatedAt() int {
	return sd.created_at
}

func (sd *movementStockDomain) SetID(id string) {
	sd.iD = id
}
