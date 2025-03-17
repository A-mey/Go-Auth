package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sv *SchemaValidatorMiddleware) ValidationRequestMiddleware(schemaMap map[string]string) gin.HandlerFunc {
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
		isValid, validationErrorMessage, err := sv.schemaValidatorService.ValidateSchema(key, string(body), schemaMap)
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
