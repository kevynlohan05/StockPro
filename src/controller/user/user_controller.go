package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/model/user/service"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	LoginUser(c *gin.Context)
	CreateUser(c *gin.Context)
	FindUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
