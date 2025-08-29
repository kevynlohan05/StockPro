package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/controller/model/response"
	"github.com/kevynlohan05/StockPro/src/view"
)

func (sc *stockControllerInterface) FindMovementsStock(c *gin.Context) {
	log.Println("Iniciando FindMovementsStock controller!")

	movementsDomain, err := sc.service.FindMovementsStock()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if len(movementsDomain) == 0 {
		errorMessage := rest_err.NewNotFoundError("Não foi encontrado movimentações no estoque.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var movementsResponse []response.StockMovementResponse

	for _, movement := range movementsDomain {
		movementsResponse = append(movementsResponse, view.ConvertMovementStockDomainToResponse(movement))
	}

	c.JSON(http.StatusOK, movementsResponse)
}
