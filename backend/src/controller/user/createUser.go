package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/configuration/validation"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
	"github.com/kevynlohan05/StockPro/src/view"
)

var (
	UserDomainInterface userModel.UserDomainInterface
)

func (ud *userControllerInterface) CreateUser(c *gin.Context) {
	log.Println("Iniciando CreateUser controller")

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := userModel.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
		userRequest.Sector,
		userRequest.Role,
		true,
	)

	domainResult, err := ud.service.CreateUserServices(domain)
	if err != nil {
		log.Println("Erro ao criar usuário")
		c.JSON(err.Code, err)
	}

	if domainResult == nil {
		log.Println("Erro: domainResult é nil")
	}

	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(domainResult))
}
