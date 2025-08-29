package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/model/product/service"
)

func NewProductControllerInterface(serviceInterface service.ProductDomainService) ProductControllerInterface {
	return &productControllerInterface{
		service: serviceInterface,
	}
}

type ProductControllerInterface interface {
	CreateProduct(c *gin.Context)
	FindProduct(c *gin.Context)
	FindAllProducts(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type productControllerInterface struct {
	service service.ProductDomainService
}
