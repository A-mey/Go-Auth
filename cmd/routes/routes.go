package routes

import (
	healthRoutes "github.com/A-mey/GO-AUTH/api/health/routes"
	userRoutes "github.com/A-mey/GO-AUTH/api/v1/users/routes"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	user := &userRoutes.UserRoutes{}
	health := &healthRoutes.HealthRoutes{}

	user.InitializeRoutes(r)
	health.InitializeRoutes(r)
}
