package response

type StockMovementResponse struct {
	ID        string `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    int    `json:"user_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	Reason    string `json:"reason"`
	CreatedAt int    `json:"created_at"`
}
