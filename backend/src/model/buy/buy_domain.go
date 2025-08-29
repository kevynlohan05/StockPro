package model

type buyDomain struct {
	iD          string
	userID      int
	payment     string
	items       []BuyItemDomainInterface
	observation string
	total       float64
}

type buyItemDomain struct {
	iD        string
	productID int
	quantity  int
	price     float64
}

func (b *buyDomain) GetID() string {
	return b.iD
}

func (b *buyDomain) GetUserID() int {
	return b.userID
}

func (b *buyDomain) GetPaymentMethod() string {
	return b.payment
}

func (b *buyDomain) GetItems() []BuyItemDomainInterface {
	return b.items
}

func (b *buyDomain) GetTotal() float64 {
	return b.total
}

func (b *buyDomain) GetObservation() string {
	return b.observation
}

func (b *buyDomain) SetID(id string) {
	b.iD = id
}

func (b *buyItemDomain) GetProductID() int {
	return b.productID
}

func (b *buyItemDomain) GetQuantity() int {
	return b.quantity
}

func (b *buyItemDomain) GetPrice() float64 {
	return b.price
}

func (b *buyItemDomain) SetID(id string) {
	b.iD = id
}
