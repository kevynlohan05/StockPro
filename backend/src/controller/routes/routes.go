package routes

import (
	"github.com/gin-gonic/gin"
	controllerBuy "github.com/kevynlohan05/StockPro/src/controller/buy"
	controllerProduct "github.com/kevynlohan05/StockPro/src/controller/product"
	controllerSale "github.com/kevynlohan05/StockPro/src/controller/sale"
	controllerStock "github.com/kevynlohan05/StockPro/src/controller/stock"
	controllerUser "github.com/kevynlohan05/StockPro/src/controller/user"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

// InitRoutes registers all application routes with proper middleware and handlers
func InitRoutes(
	r *gin.RouterGroup,
	userController controllerUser.UserControllerInterface,
	productController controllerProduct.ProductControllerInterface,
	stockController controllerStock.StockControllerInterface,
	saleControllert controllerSale.SaleControllerInterface,
	buyController controllerBuy.BuyControllerInterface,
) {
	// --- User routes ---
	r.POST("/user/login", userController.LoginUser)
	r.POST("/user/createUser", userController.CreateUser)
	r.GET("/user/getUserById/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.FindUser)
	r.GET("/user/getAllUsers", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.FindAllUsers)
	r.PUT("/user/updateUser/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.UpdateUser)
	r.DELETE("/user/deleteUser/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.DeleteUser)

	// --- Product routes ---
	r.POST("/product/createProduct", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, productController.CreateProduct)
	r.GET("/product/getProductById/:productId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, productController.FindProduct)
	r.GET("/product/getAllProducts", userModel.VerifyTokenMiddleware, productController.FindAllProducts)
	r.PUT("/product/updateProduct/:productId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, productController.UpdateProduct)
	r.DELETE("/product/deleteProduct/:productId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, productController.DeleteProduct)

	// --- Stock routes ---
	r.POST("/stock/createStockMovement", userModel.VerifyTokenMiddleware, stockController.CreateStockMovement)
	r.GET("/stock/getAllMovements", userModel.VerifyTokenMiddleware, stockController.FindMovementsStock)
	r.GET("/stock/getStock", userModel.VerifyTokenMiddleware, stockController.FindStock)

	// --- Sale routes ---
	r.POST("/sale/createSale", userModel.VerifyTokenMiddleware, saleControllert.CreateSale)

	// --- Buy routes ---
	r.POST("/buy/createBuy", userModel.VerifyTokenMiddleware, buyController.CreateBuy)
}
