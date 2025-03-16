package routes

import (
	"github.com/A-mey/GO-AUTH/api/middlewares"
	RoutesInterface "github.com/A-mey/GO-AUTH/common/interfaces"
	"github.com/gin-gonic/gin"
)

var _ RoutesInterface.RoutesInterface = (*otpRoutes)(nil)

type otpRoutes struct{}

func (hr *otpRoutes) InitializeRoutes(r *gin.Engine) {

	r.Use(middlewares.ValidationMiddleware(map[string]string{"name": "John Doe", "age": "30", "email": "john@example.com"}))

	// healthService := &services.DefaultHealthService{}

	// var healthController healthInterfaces.HealthController = controllers.NewHealthController(healthService)
	userRoutes := r.Group("/otp")
	{
		userRoutes.POST("/registration")
	}
}
