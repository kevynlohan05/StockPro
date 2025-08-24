package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
)

func (pd *productDomainService) DeleteProductServices(productId string) *rest_err.RestErr {
	log.Println("Iniciando DeleteProductServices service")

	err := pd.productRepository.DeleteProduct(productId)
	if err != nil {
		log.Println("Erro ao deletar produto no banco de dados: ", err)
		return err
	}
	return nil
}
