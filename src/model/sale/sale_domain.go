package model

type saleItemDomain struct {
	iD        string
	productID int
	quantity  int
	price     float64
}

type saleDomain struct {
	iD            string
	userID        int
	paymentMethod string
	items         []SaleItemDomainInterface
	total         float64
}

func (s *saleDomain) GetUserID() int {
	return s.userID
}
func (s *saleDomain) GetPaymentMethod() string {
	return s.paymentMethod
}
func (s *saleDomain) GetItems() []SaleItemDomainInterface {
	return s.items
}
func (s *saleDomain) GetTotal() float64 {
	return s.total
}

func (s *saleDomain) SetID(id string) {
	s.iD = id
}

func (s *saleItemDomain) GetProductID() int {
	return s.productID
}

func (s *saleItemDomain) GetQuantity() int {
	return s.quantity
}

func (s *saleItemDomain) GetPrice() float64 {
	return s.price
}

func (s *saleItemDomain) SetID(id string) {
	s.iD = id
}
