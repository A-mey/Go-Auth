package interfaces

import "github.com/gin-gonic/gin"

type RoutesInterface interface {
	InitializeRoutes(r *gin.Engine)
}
