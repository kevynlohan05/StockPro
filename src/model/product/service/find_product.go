package service

import (
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pd *productDomainService) FindProductByIdServices(id string) (productModel.ProductDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
