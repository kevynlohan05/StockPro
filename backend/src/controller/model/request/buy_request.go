package request

type BuyRequest struct {
	UserID        int              `json:"user_id" binding:"required"`
	PaymentMethod string           `json:"payment_method" binding:"required"`
	Items         []BuyItemRequest `json:"items" binding:"required,dive,required"`
	Observation   string           `json:"observation"`
}

type BuyItemRequest struct {
	ProductID int     `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gt=0"`
}
