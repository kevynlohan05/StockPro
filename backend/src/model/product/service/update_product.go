package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pd *productDomainService) UpdateProduct(productId string, productDomain productModel.ProductDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando UpdateProduct service")

	err := pd.productRepository.UpdateProduct(productId, productDomain)
	if err != nil {
		log.Println("Erro ao atualizar produto no banco de dados: ", err)
		return err
	}
	return nil
}
