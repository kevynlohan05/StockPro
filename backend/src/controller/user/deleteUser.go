package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	log.Println("Iniciando DeleteUser controller")

	userId := c.Param("userId")

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		log.Println("Erro ao deletar usuário: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
