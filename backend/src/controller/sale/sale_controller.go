package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/model/sale/service"
)

func NewSaleControllerInterface(serviceInterface service.SaleDomainService) SaleControllerInterface {
	return &saleControllerInterface{
		service: serviceInterface,
	}
}

type saleControllerInterface struct {
	service service.SaleDomainService
}

type SaleControllerInterface interface {
	CreateSale(c *gin.Context)
}
