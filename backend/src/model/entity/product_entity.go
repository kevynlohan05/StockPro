package entity

type ProductEntity struct {
	ID            int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Mark          string  `json:"mark"`
	PurchasePrice float64 `json:"purchasePrice"`
	SalePrice     float64 `json:"salePrice"`
	Images        string  `json:"image"`
}
