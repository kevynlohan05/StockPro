package model

type ProductDomainInterface interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetMark() string
	GetPurchasePrice() string
	GetSalePrice() string
	GetImages() []string

	SetID(id string)
}

func NewProductDomainService(name, description, mark, purchasePrice, salePrice string, images []string) ProductDomainInterface {
	return &productDomain{
		name:          name,
		description:   description,
		mark:          mark,
		purchasePrice: purchasePrice,
		salePrice:     salePrice,
		images:        images,
	}
}

func NewProductUpdateDomainService(name, description, mark, purchasePrice, salePrice string) ProductDomainInterface {
	return &productDomain{
		name:          name,
		description:   description,
		mark:          mark,
		purchasePrice: purchasePrice,
		salePrice:     salePrice,
	}
}
