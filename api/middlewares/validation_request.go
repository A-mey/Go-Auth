package middlewares

import (
	"net/http"

	"github.com/A-mey/GO-AUTH/libs/validator"
	"github.com/gin-gonic/gin"
)

func ValidationRequestMiddleware(schemaMap map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValid, validationErrorMessage, err := validator.Validator(c, schemaMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			c.Abort()
			return
		}
		if !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": validationErrorMessage})
			c.Abort()
			return
		}
		c.Next()
	}
}
