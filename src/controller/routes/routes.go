package routes

import (
	"github.com/gin-gonic/gin"
	controllerUser "github.com/kevynlohan05/StockPro/src/controller/user"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

// InitRoutes registers all application routes with proper middleware and handlers
func InitRoutes(
	r *gin.RouterGroup,
	userController controllerUser.UserControllerInterface,
) {
	// --- User routes ---
	r.POST("/user/login", userController.LoginUser)

	r.POST("/user/createUser", userController.CreateUser)
	r.GET("/user/getUserById/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.FindUser)
	r.PUT("/user/updateUser/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.UpdateUser)
	r.DELETE("/user/deleteUser/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.DeleteUser)

}
