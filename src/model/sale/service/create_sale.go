package service

import (
	"log"
	"strconv"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	saleModel "github.com/kevynlohan05/StockPro/src/model/sale"
)

func (s *saleDomainService) CreateSaleService(sale saleModel.SaleDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando CreateSaleService")

	// 1️⃣ Checa estoque de cada item
	for _, item := range sale.GetItems() {
		ok, err := s.repositoryStock.CheckStock(item.GetProductID(), item.GetQuantity())
		if err != nil {
			return err
		}
		if !ok {
			return rest_err.NewBadRequestError("Estoque insuficiente para produto ID " + strconv.Itoa(item.GetProductID()))
		}
	}

	if err := s.repositorySale.CreateSaleTransaction(sale); err != nil {
		return err
	}

	return nil
}
