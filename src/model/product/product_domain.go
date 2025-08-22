package model

type productDomain struct {
	iD            string
	name          string
	description   string
	mark          string
	purchasePrice string
	salePrice     string
	image         string
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

func (pd *productDomain) GetpurchasePrice() string {
	return pd.purchasePrice
}

func (pd *productDomain) GetsalePrice() string {
	return pd.salePrice
}

func (pd *productDomain) Getimage() string {
	return pd.image
}
