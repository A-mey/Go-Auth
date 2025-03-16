package healthRoutes

import (
	"github.com/A-mey/GO-AUTH/api/health/controllers"
	healthInterfaces "github.com/A-mey/GO-AUTH/api/health/interfaces"
	"github.com/A-mey/GO-AUTH/api/health/services"
	RoutesInterface "github.com/A-mey/GO-AUTH/common/interfaces"
	"github.com/gin-gonic/gin"
)

var _ RoutesInterface.RoutesInterface = (*HealthRoutes)(nil)

type HealthRoutes struct{}

func (hr *HealthRoutes) InitializeRoutes(r *gin.Engine) {

	healthService := &services.DefaultHealthService{}

	var healthController healthInterfaces.HealthController = controllers.NewHealthController(healthService)
	userRoutes := r.Group("/health")
	{
		userRoutes.GET("/", healthController.HealthCheck)
	}
}
