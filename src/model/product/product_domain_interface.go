package model

type ProductDomainInterface interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetMark() string
	GetpurchasePrice() string
	GetsalePrice() string
	Getimage() string
}

func NewProductDomainService(name, description, mark, purchasePrice, salePrice, image string) ProductDomainInterface {
	return &productDomain{
		name:          name,
		description:   description,
		mark:          mark,
		purchasePrice: purchasePrice,
		salePrice:     salePrice,
		image:         image,
	}
}

func NewProductUpdateDomainService(name, description, mark, purchasePrice, salePrice, image string) ProductDomainInterface {
	return &productDomain{
		name:          name,
		description:   description,
		mark:          mark,
		purchasePrice: purchasePrice,
		salePrice:     salePrice,
		image:         image,
	}
}
