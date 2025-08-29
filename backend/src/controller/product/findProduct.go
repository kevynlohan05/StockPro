package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/view"
)

func (pc *productControllerInterface) FindProduct(c *gin.Context) {
	log.Println("Iniciando FindProduct controller")

	productId := c.Param("productId")

	domainResult, err := pc.service.FindProductByIdServices(productId)
	if err != nil {
		log.Println("Erro ao buscar produto: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertProductDomainToResponse(domainResult))
}
