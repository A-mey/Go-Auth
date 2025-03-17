package middlewares

import (
	"net/http"

	"github.com/A-mey/GO-AUTH/libs/validator"
	"github.com/gin-gonic/gin"
)

func ValidationRequestMiddleware(schemaMap map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// isValid, validationErrorMessage, err := validator.Validator(c, schemaMap)
		originalRoute := c.FullPath()
		method := c.Request.Method
		key := originalRoute + "/" + method
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			c.Abort()
			return
		}
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		// 	c.Abort()
		// 	return
		// }
		isValid, validationErrorMessage, err := validator.Validator(c, schemaMap)
		if !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": validationErrorMessage})
			c.Abort()
			return
		}
		c.Next()
	}
}
