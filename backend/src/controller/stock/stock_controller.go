package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/model/stock/service"
)

func NewStockControllerInterface(serviceInterface service.StockDomainService) StockControllerInterface {
	return &stockControllerInterface{
		service: serviceInterface,
	}
}

type StockControllerInterface interface {
	CreateStockMovement(c *gin.Context)
}

type stockControllerInterface struct {
	service service.StockDomainService
}
