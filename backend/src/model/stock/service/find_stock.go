package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
)

func (s *stockDomainService) FindMovementsStock() ([]stockModel.StockDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindMovementsStock service!")

	movements, err := s.stockRepository.FindMovementsStock()
	if err != nil {
		return nil, err
	}

	return movements, nil
}

func (s *stockDomainService) FindStock() ([]stockModel.StockQuantityDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindMovementsStock service!")

	stock, err := s.stockRepository.FindStockQuantity()
	if err != nil {
		return nil, err
	}

	return stock, nil
}
