package entity

type StockEntity struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID int    `json:"product_id"`
	UserID    int    `json:"user_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	Reason    string `json:"reason"`
}
