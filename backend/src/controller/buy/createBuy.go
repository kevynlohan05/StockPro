package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	buyModel "github.com/kevynlohan05/StockPro/src/model/buy"
)

func (bc *buyControllerInterface) CreateBuy(c *gin.Context) {
	log.Println("Iniciando CreateBuy controller")

	var buyRequest request.BuyRequest
	if err := c.ShouldBindJSON(&buyRequest); err != nil {
		log.Println("Erro ao fazer bind do JSON:", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var items []buyModel.BuyItemDomainInterface
	for _, itemReq := range buyRequest.Items {
		itemDomain := buyModel.NewBuyDomainItem(itemReq.ProductID, itemReq.Quantity, itemReq.Price)
		items = append(items, itemDomain)
	}

	domain := buyModel.NewBuyDomain(buyRequest.UserID, buyRequest.PaymentMethod, items, buyRequest.Observation)

	if err := bc.service.CreateBuyService(domain); err != nil {
		log.Println("Erro ao criar compra:", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Compra registrada com sucesso",
	})
}
