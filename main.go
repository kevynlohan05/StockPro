package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kevynlohan05/StockPro/src/configuration/database/mysql"
	controllerProduct "github.com/kevynlohan05/StockPro/src/controller/product"
	"github.com/kevynlohan05/StockPro/src/controller/routes"
	controllerStock "github.com/kevynlohan05/StockPro/src/controller/stock"
	controllerUser "github.com/kevynlohan05/StockPro/src/controller/user"
	repositoryProduct "github.com/kevynlohan05/StockPro/src/model/product/repository"
	serviceProduct "github.com/kevynlohan05/StockPro/src/model/product/service"
	repositoryStock "github.com/kevynlohan05/StockPro/src/model/stock/repository"
	serviceStock "github.com/kevynlohan05/StockPro/src/model/stock/service"
	repositoryUser "github.com/kevynlohan05/StockPro/src/model/user/repository"
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
	routes.InitRoutes(&router.RouterGroup, userController, productController, stockController)

	// Inicia servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
