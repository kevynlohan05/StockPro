package repository

import (
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pr *productRepository) UpdateProduct(productId string, productDomain productModel.ProductDomainInterface) *rest_err.RestErr {
	return nil
}
