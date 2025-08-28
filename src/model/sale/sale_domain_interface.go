package model

type SaleDomainInterface interface {
	GetUserID() int
	GetPaymentMethod() string
	GetItems() []SaleItemDomainInterface
	GetTotal() float64
	SetID(id string)
}

type SaleItemDomainInterface interface {
	GetProductID() int
	GetQuantity() int
	GetPrice() float64
	SetID(id string)
}

func NewSaleDomain(userID int, paymentMethod string, items []SaleItemDomainInterface) SaleDomainInterface {
	var total float64
	for _, item := range items {
		total += float64(item.GetQuantity()) * item.GetPrice()
	}

	return &saleDomain{
		userID:        userID,
		paymentMethod: paymentMethod,
		items:         items,
		total:         total,
	}
}

func NewSaleDomainItem(productID int, quantity int, price float64) SaleItemDomainInterface {
	return &saleItemDomain{
		productID: productID,
		quantity:  quantity,
		price:     price,
	}
}
