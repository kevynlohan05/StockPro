package routes

import (
	"github.com/gin-gonic/gin"
	controllerProduct "github.com/kevynlohan05/StockPro/src/controller/product"
	controllerUser "github.com/kevynlohan05/StockPro/src/controller/user"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

// InitRoutes registers all application routes with proper middleware and handlers
func InitRoutes(
	r *gin.RouterGroup,
	userController controllerUser.UserControllerInterface,
	productController controllerProduct.ProductControllerInterface,
) {
	// --- User routes ---
	r.POST("/user/login", userController.LoginUser)

	r.POST("/user/createUser", userController.CreateUser)
	r.GET("/user/getUserById/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.FindUser)
	r.PUT("/user/updateUser/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.UpdateUser)
	r.DELETE("/user/deleteUser/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.DeleteUser)

	r.POST("/product/createProduct", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, productController.CreateProduct)
	r.GET("/product/getProductById/:productId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, productController.FindProduct)
	r.PUT("/product/updateProduct/:productId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, productController.UpdateProduct)
	r.DELETE("/product/deleteProduct/:productId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, productController.DeleteProduct)

}
