package entity

type ProductEntity struct {
	ID            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Mark          string `json:"mark"`
	PurchasePrice string `json:"purchasePrice"`
	SalePrice     string `json:"salePrice"`
	Image         string `json:"image"`
}
