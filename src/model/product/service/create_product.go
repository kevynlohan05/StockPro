package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pd *productDomainService) CreateProductServices(productDomain productModel.ProductDomainInterface) (productModel.ProductDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando CreateProductServices service")

	productResult, err := pd.productRepository.CreateProduct(productDomain)
	if err != nil {
		log.Println("Erro ao criar produto no banco de dados: ", err)
		return nil, err
	}

	return productResult, nil
}
