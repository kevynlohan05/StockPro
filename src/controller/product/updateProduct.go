package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/configuration/validation"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pc *productControllerInterface) UpdateProduct(c *gin.Context) {
	log.Println("Iniciando UpdateProduct controller")

	productId := c.Param("productId")

	var productRequest request.ProductRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := productModel.NewProductUpdateDomainService(
		productRequest.Name,
		productRequest.Description,
		productRequest.Mark,
		productRequest.PurchasePrice,
		productRequest.SalePrice,
		productRequest.Image,
	)

	err := pc.service.UpdateProduct(productId, domain)
	if err != nil {
		log.Println("Erro ao atualizar produto: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
