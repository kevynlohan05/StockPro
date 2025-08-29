package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/configuration/validation"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	log.Println("Iniciando UpdateUser controller")

	var userResquest request.UserUpdateRequest

	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := userModel.NewUserUpdateDomain(
		userResquest.Name,
		userResquest.Email,
		userResquest.Password,
		userResquest.Sector,
		userResquest.Role,
		userResquest.Active,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		log.Println("Erro ao atualizar usuário")
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)

}
