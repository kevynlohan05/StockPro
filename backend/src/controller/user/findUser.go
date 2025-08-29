package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/view"
)

func (uc *userControllerInterface) FindUser(c *gin.Context) {
	log.Println("Iniciando FindUser controller")

	userId := c.Param("userId")

	user, err := uc.service.FindUserByIdServices(userId)
	if err != nil {
		log.Println("Erro ao buscar usuário: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(user))
}
