package request

type StockMovementRequest struct {
	ProductID int    `json:"product_id" binding:"required"`
	UserID    int    `json:"user_id" binding:"required"`
	Type      string `json:"type" binding:"required,oneof=entrada saida"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
	Reason    string `json:"reason"`
}
