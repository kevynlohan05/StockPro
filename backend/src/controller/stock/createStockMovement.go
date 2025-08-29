package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/configuration/validation"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
)

func (sc *stockControllerInterface) CreateStockMovement(c *gin.Context) {
	log.Println("Iniciando CreateStockMovement controller!")

	var stockRequest request.StockMovementRequest
	if err := c.ShouldBindJSON(&stockRequest); err != nil {
		errRest := validation.ValidateRequestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	stockDomain := stockModel.NewStockDomain(
		stockRequest.ProductID,
		stockRequest.UserID,
		stockRequest.Type,
		stockRequest.Quantity,
		stockRequest.Reason,
	)

	err := sc.service.CreateStockMovementService(stockDomain)
	if err != nil {
		log.Println("Erro ao criar movimentação de estoque: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movimentação registrada com sucesso!"})
}
