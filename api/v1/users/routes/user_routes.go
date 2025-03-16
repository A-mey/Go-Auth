package userRoutes

import (
	"github.com/A-mey/GO-AUTH/api/v1/users/controllers"
	"github.com/A-mey/GO-AUTH/api/v1/users/interfaces"
	"github.com/A-mey/GO-AUTH/api/v1/users/services"
	RoutesInterface "github.com/A-mey/GO-AUTH/common/interfaces"
	"github.com/gin-gonic/gin"
)

var _ RoutesInterface.RoutesInterface = (*UserRoutes)(nil)

type UserRoutes struct{}

func (s *UserRoutes) InitializeRoutes(r *gin.Engine) {

	userService := &services.DefaultUserService{}

	var userController interfaces.UserControllerInterface = controllers.NewUserController(userService)
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userController.CreateNewUserController)
		userRoutes.GET("/", userController.DeleteUserController)
	}
}
