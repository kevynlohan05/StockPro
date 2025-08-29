package controller

import (
	"github.com/kevynlohan05/StockPro/src/model/buy/service"
)

func NewBuyControllerInterface(serviceInterface service.BuyDomainService) BuyControllerInterface {
	return &buyControllerInterface{
		service: serviceInterface,
	}
}

type buyControllerInterface struct {
	service service.BuyDomainService
}

type BuyControllerInterface interface {
}
