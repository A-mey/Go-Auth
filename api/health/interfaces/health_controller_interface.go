package interfaces

import "github.com/gin-gonic/gin"

type HealthController interface {
	HealthCheck(c *gin.Context)
}
