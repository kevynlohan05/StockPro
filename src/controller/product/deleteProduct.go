package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pc *productControllerInterface) DeleteProduct(c *gin.Context) {
	log.Println("Iniciando DeleteProduct controller")

	productId := c.Param("productId")

	err := pc.service.DeleteProductServices(productId)
	if err != nil {
		log.Println("Erro ao deletar produto: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
