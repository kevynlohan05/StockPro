package controller

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kevynlohan05/StockPro/src/controller/model/request"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
	"github.com/kevynlohan05/StockPro/src/view"
)

var validate = validator.New()

func (pc *productControllerInterface) CreateProduct(c *gin.Context) {
	log.Println("Iniciando CreateProduct controller")

	if c.ContentType() != "multipart/form-data" {
		log.Println("Invalid content type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be multipart/form-data"})
		return
	}

	productRequest := request.ProductRequest{
		Name:          c.PostForm("name"),
		Description:   c.PostForm("description"),
		Mark:          c.PostForm("mark"),
		PurchasePrice: c.PostForm("purchase_price"),
		SalePrice:     c.PostForm("sale_price"),
	}

	if err := validate.Struct(productRequest); err != nil {
		log.Println("Validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	var images []string

	form, _ := c.MultipartForm()
	if form != nil && form.File != nil {
		files := form.File["images"]
		for _, file := range files {
			contentType := file.Header.Get("Content-Type")
			if !strings.HasPrefix(contentType, "image/") {
				log.Println("Invalid file type:", contentType)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Only image files are allowed"})
				return
			}

			if _, err := os.Stat("uploads"); os.IsNotExist(err) {
				err = os.MkdirAll("uploads", 0755)
				if err != nil {
					log.Println("Failed to create uploads directory:", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare uploads folder"})
					return
				}
			}

			path := "uploads/" + file.Filename
			if err := c.SaveUploadedFile(file, path); err != nil {
				log.Println("Failed to save file:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
				return
			}

			baseURL := os.Getenv("BASE_URL")
			publicURL := baseURL + "/" + path
			images = append(images, publicURL)
		}
	}

	domain := productModel.NewProductDomainService(
		productRequest.Name,
		productRequest.Description,
		productRequest.Mark,
		productRequest.PurchasePrice,
		productRequest.SalePrice,
		images,
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
