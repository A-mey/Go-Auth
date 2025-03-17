package middlewares

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func ValidateUserMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Authorization header required"})
// 			c.Abort()
// 			return
// 		}
// 	}
// }
