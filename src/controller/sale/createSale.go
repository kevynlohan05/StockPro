package controller

import (
	"log"
	"net/http"

	"github.com/kevynlohan05/StockPro/src/configuration/validation"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	saleModel "github.com/kevynlohan05/StockPro/src/model/sale"

	"github.com/gin-gonic/gin"
)

func (sc *saleControllerInterface) CreateSale(c *gin.Context) {
	log.Println("Iniciando CreateSale controller")

	var saleRequest request.SaleRequest
	if err := c.ShouldBindJSON(&saleRequest); err != nil {
		errRest := validation.ValidateRequestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	var items []saleModel.SaleItemDomainInterface
	for _, itemReq := range saleRequest.Items {
		itemDomain := saleModel.NewSaleDomainItem(itemReq.ProductID, itemReq.Quantity, itemReq.Price)
		items = append(items, itemDomain)
	}

	domain := saleModel.NewSaleDomain(saleRequest.UserID, saleRequest.PaymentMethod, items)

	if err := sc.service.CreateSaleService(domain); err != nil {
		log.Println("Erro ao criar venda:", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Venda registrada com sucesso",
	})
}
