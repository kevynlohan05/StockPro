package request

type ProductRequest struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Mark          string `json:"mark" binding:"required"`
	PurchasePrice string `json:"purchase_price" binding:"required"`
	SalePrice     string `json:"sale_price" binding:"required"`
}

type ProductUpdateRequest struct {
	Name          string `json:"name" binding:"omitempty"`
	Description   string `json:"description" binding:"omitempty"`
	Mark          string `json:"mark" binding:"omitempty"`
	PurchasePrice string `json:"purchase_price" binding:"omitempty"`
	SalePrice     string `json:"sale_price" binding:"omitempty"`
}
