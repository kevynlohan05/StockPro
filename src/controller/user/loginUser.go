package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/configuration/validation"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ud *userControllerInterface) LoginUser(c *gin.Context) {
	log.Println("Iniciando LoginUser controller")

	var userLoginRequest request.LoginRequest

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := userModel.NewLoginDomain(
		userLoginRequest.Email,
		userLoginRequest.Password,
	)

	token, err := ud.service.LoginUserServices(domain)
	if err != nil {
		log.Println("Erro ao realizar login do usuário")
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
