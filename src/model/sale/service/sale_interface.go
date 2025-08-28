package service

import (
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	saleModel "github.com/kevynlohan05/StockPro/src/model/sale"
	repositorySale "github.com/kevynlohan05/StockPro/src/model/sale/repository"
	repositoryStock "github.com/kevynlohan05/StockPro/src/model/stock/repository"
)

func NewSaleDomainService(saleRepository repositorySale.SaleRepository, stockRepository repositoryStock.StockRepository) SaleDomainService {
	return &saleDomainService{saleRepository, stockRepository}
}

type saleDomainService struct {
	repositorySale  repositorySale.SaleRepository
	repositoryStock repositoryStock.StockRepository
}

type SaleDomainService interface {
	CreateSaleService(sale saleModel.SaleDomainInterface) *rest_err.RestErr
}
