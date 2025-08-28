package service

import (
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
	repositoryStock "github.com/kevynlohan05/StockPro/src/model/stock/repository"
)

func NewStockDomainService(stockRepository repositoryStock.StockRepository) StockDomainService {
	return &stockDomainService{stockRepository}
}

type stockDomainService struct {
	stockRepository repositoryStock.StockRepository
}

type StockDomainService interface {
	CreateStockMovementService(stockDomain stockModel.StockDomainInterface) *rest_err.RestErr
}
