package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pd *productDomainService) FindProductByIdServices(id string) (productModel.ProductDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindProductByIdServices service")

	product, err := pd.productRepository.FindProductById(id)
	if err != nil {
		log.Println("Erro ao buscar produto no banco de dados: ", err)
		return nil, err
	}

	return product, nil
}

func (pd *productDomainService) FindAllProductsServices() ([]productModel.ProductDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindAllProductsServices service")

	products, err := pd.productRepository.FindAllProducts()
	if err != nil {
		log.Println("Erro ao buscar produtos no banco de dados: ", err)
		return nil, err
	}

	return products, nil
}
