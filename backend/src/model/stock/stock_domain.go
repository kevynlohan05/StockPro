package model

type StockDomain struct {
	iD         string
	productID  int
	userID     int
	stockType  string
	quantity   int
	reason     string
	created_at int
}

func (sd *StockDomain) GetID() string {
	return sd.iD
}

func (sd *StockDomain) GetProductID() int {
	return sd.productID
}

func (sd *StockDomain) GetUserID() int {
	return sd.userID
}

func (sd *StockDomain) GetStockType() string {
	return sd.stockType
}

func (sd *StockDomain) GetQuantity() int {
	return sd.quantity
}

func (sd *StockDomain) GetReason() string {
	return sd.reason
}

func (sd *StockDomain) GetCreatedAt() int {
	return sd.created_at
}

func (sd *StockDomain) SetID(id string) {
	sd.iD = id
}
