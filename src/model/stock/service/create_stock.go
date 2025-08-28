package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
)

func (s *stockDomainService) CreateStockMovementService(stockDomain stockModel.StockDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando CreateStockMovementService service")

	if stockDomain.GetStockType() == "saida" {
		ok, err := s.stockRepository.CheckStock(stockDomain.GetProductID(), stockDomain.GetQuantity())
		if err != nil {
			return err
		}
		if !ok {
			return rest_err.NewBadRequestError("Estoque insuficiente para saída")
		}
	}

	if err := s.stockRepository.CreateStockMovement(stockDomain); err != nil {
		return err
	}

	return nil
}
