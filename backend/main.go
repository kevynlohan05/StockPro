package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kevynlohan05/StockPro/src/configuration/database/mysql"
	controllerBuy "github.com/kevynlohan05/StockPro/src/controller/buy"
	controllerProduct "github.com/kevynlohan05/StockPro/src/controller/product"
	"github.com/kevynlohan05/StockPro/src/controller/routes"
	controllerSale "github.com/kevynlohan05/StockPro/src/controller/sale"
	controllerStock "github.com/kevynlohan05/StockPro/src/controller/stock"
	controllerUser "github.com/kevynlohan05/StockPro/src/controller/user"

	repositoryBuy "github.com/kevynlohan05/StockPro/src/model/buy/repository"
	repositoryProduct "github.com/kevynlohan05/StockPro/src/model/product/repository"
	repositorySale "github.com/kevynlohan05/StockPro/src/model/sale/repository"
	repositoryStock "github.com/kevynlohan05/StockPro/src/model/stock/repository"
	repositoryUser "github.com/kevynlohan05/StockPro/src/model/user/repository"

	serviceBuy "github.com/kevynlohan05/StockPro/src/model/buy/service"
	serviceProduct "github.com/kevynlohan05/StockPro/src/model/product/service"
	serviceSale "github.com/kevynlohan05/StockPro/src/model/sale/service"
	serviceStock "github.com/kevynlohan05/StockPro/src/model/stock/service"
	serviceUser "github.com/kevynlohan05/StockPro/src/model/user/service"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Conexão MySQL
	database, err := mysql.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to MySQL, error=%s \n", err.Error())
		return
	}

	// Inicializa repositórios e serviços com *sql.DB
	repoUser := repositoryUser.NewUserRepository(database)
	userServiceInstance := serviceUser.NewUserDomainService(repoUser)
	userController := controllerUser.NewUserControllerInterface(userServiceInstance)

	repoProduct := repositoryProduct.NewProductRepository(database)
	productServiceInstance := serviceProduct.NewProductDomainService(repoProduct)
	productController := controllerProduct.NewProductControllerInterface(productServiceInstance)

	repoStock := repositoryStock.NewStockRepository(database)
	stockServiceInstance := serviceStock.NewStockDomainService(repoStock)
	stockController := controllerStock.NewStockControllerInterface(stockServiceInstance)

	repoSale := repositorySale.NewSaleRepository(database)
	saleServiceInstance := serviceSale.NewSaleDomainService(repoSale, repoStock)
	saleController := controllerSale.NewSaleControllerInterface(saleServiceInstance)

	repoBuy := repositoryBuy.NewBuyRepository(database)
	buyServiceInstance := serviceBuy.NewBuyDomainService(repoBuy, repoStock)
	buyController := controllerBuy.NewBuyControllerInterface(buyServiceInstance)

	// Setup Gin com CORS
	router := gin.Default()

	router.Static("/uploads", "./uploads")

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173", // Frontend local
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Inicializa rotas
	routes.InitRoutes(&router.RouterGroup, userController, productController, stockController, saleController, buyController)

	// Inicia servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
