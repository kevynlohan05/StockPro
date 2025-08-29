package service

import (
	repositoryBuy "github.com/kevynlohan05/StockPro/src/model/buy/repository"
	repositoryStock "github.com/kevynlohan05/StockPro/src/model/stock/repository"
)

func NewBuyDomainService(buyRepository repositoryBuy.BuyRepository, stockRepository repositoryStock.StockRepository) BuyDomainService {
	return &buyDomainService{buyRepository, stockRepository}
}

type buyDomainService struct {
	repositoryBuy   repositoryBuy.BuyRepository
	repositoryStock repositoryStock.StockRepository
}

type BuyDomainService interface {
}
