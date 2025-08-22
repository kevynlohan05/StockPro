package controller

import "github.com/gin-gonic/gin"

func NewProductControllerInterface(serviceInterface service.ProductDomainService) ProductControllerInterface {
	return &productControllerInterface{
		service: serviceInterface,
	}
}

type ProductControllerInterface interface {
	CreateProduct(c *gin.Context)
	FindProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type productControllerInterface struct {
	service service.ProductDomainService
}
