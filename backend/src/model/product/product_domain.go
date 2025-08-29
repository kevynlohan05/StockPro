package model

type productDomain struct {
	iD            string
	name          string
	description   string
	mark          string
	purchasePrice string
	salePrice     string
	images        []string
}

func (pd *productDomain) GetID() string {
	return pd.iD
}

func (pd *productDomain) GetName() string {
	return pd.name
}

func (pd *productDomain) GetDescription() string {
	return pd.description
}

func (pd *productDomain) GetMark() string {
	return pd.mark
}

func (pd *productDomain) GetPurchasePrice() string {
	return pd.purchasePrice
}

func (pd *productDomain) GetSalePrice() string {
	return pd.salePrice
}

func (pd *productDomain) GetImages() []string {
	return pd.images
}

func (pd *productDomain) SetID(id string) {
	pd.iD = id
}
