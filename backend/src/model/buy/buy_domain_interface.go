package model

type BuyDomainInterface interface {
	GetID() string
	SetID(id string)
	GetUserID() int
	GetPaymentMethod() string
	GetItems() []BuyItemDomainInterface
	GetTotal() float64
	GetObservation() string
}

type BuyItemDomainInterface interface {
	GetProductID() int
	GetQuantity() int
	GetPrice() float64
	SetID(id string)
}

func NewBuyDomain(userID int, payment string, items []BuyItemDomainInterface, observation string) BuyDomainInterface {
	var total float64
	for _, item := range items {
		total += float64(item.GetQuantity()) * item.GetPrice()
	}
	return &buyDomain{
		userID:      userID,
		payment:     payment,
		items:       items,
		observation: observation,
		total:       total,
	}
}

func NewBuyDomainItem(productID int, quantity int, price float64) BuyItemDomainInterface {
	return &buyItemDomain{
		productID: productID,
		quantity:  quantity,
		price:     price,
	}
}
