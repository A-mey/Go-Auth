package userRoutes

import (
	"github.com/A-mey/Auth_db/api/v1/users/controllers"
	"github.com/A-mey/Auth_db/api/v1/users/interfaces"
	"github.com/A-mey/Auth_db/api/v1/users/services"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	userService := &services.DefaultUserService{}

	// userController := controllers.NewUserController(userService)
	var userController interfaces.UserControllerInterface = controllers.NewUserController(userService)
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userController.CreateNewUserController)
		userRoutes.GET("/", userController.DeleteUserController)
	}
}
