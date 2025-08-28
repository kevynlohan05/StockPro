package request

type SaleItemRequest struct {
	ProductID int     `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gt=0"`
}

type SaleRequest struct {
	UserID        int               `json:"user_id" binding:"required"`
	PaymentMethod string            `json:"payment_method" binding:"required"`
	Items         []SaleItemRequest `json:"items" binding:"required,dive,required"`
}
