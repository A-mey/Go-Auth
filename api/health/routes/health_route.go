package routes

import (
	"github.com/A-mey/Auth_db/api/health/controllers"
	"github.com/A-mey/Auth_db/api/health/interfaces"
	"github.com/A-mey/Auth_db/api/health/services"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	healthService := &services.DefaultHealthService{}

	var healthController interfaces.HealthController = controllers.NewHealthController(healthService)
	userRoutes := r.Group("/health")
	{
		userRoutes.GET("/", healthController.HealthCheck)
	}
}
