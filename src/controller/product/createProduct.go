package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/StockPro/src/configuration/validation"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
	"github.com/kevynlohan05/StockPro/src/view"
)

var (
	ProductDomainInterface productModel.ProductDomainInterface
)

func (pc *productControllerInterface) CreateProduct(c *gin.Context) {
	log.Println("Iniciando CreateProduct controller")

	var productRequest request.ProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := productModel.NewProductDomainService(
		productRequest.Name,
		productRequest.Description,
		productRequest.Mark,
		productRequest.PurchasePrice,
		productRequest.SalePrice,
		productRequest.Image,
	)

	domainResult, err := pc.service.CreateProductServices(domain)
	if err != nil {
		log.Println("Erro ao criar novo produto")
		c.JSON(err.Code, err)
		return
	}

	if domainResult == nil {
		log.Println("Erro: domainResult é nil")
		c.JSON(http.StatusInternalServerError, "Criação de produto falhou")
		return
	}

	c.JSON(http.StatusOK, view.ConvertProductDomainToResponse(domainResult))
}
