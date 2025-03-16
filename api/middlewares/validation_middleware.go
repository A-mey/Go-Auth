package middlewares

import (
	"net/http"

	"github.com/A-mey/GO-AUTH/libs/validator"
	"github.com/gin-gonic/gin"
)

// func ValidationMiddleware (map Map[string]string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		isValid, validationErrorMessage, err := validator.Validator(c, map)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		if isValid == false {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": validationErrorMessage})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}

// }

func ValidationMiddleware(schemaMap map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValid, validationErrorMessage, err := validator.Validator(c, schemaMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			c.Abort() // Stop request from proceeding
			return
		}
		if !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": validationErrorMessage})
			c.Abort()
			return
		}
		c.Next() // Proceed if validation is successful
	}
}
