package response

type ProductResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description" `
	Mark          string `json:"mark"`
	PurchasePrice string `json:"purchase_price"`
	SalePrice     string `json:"sale_price"`
	Image         string `json:"image"`
}
