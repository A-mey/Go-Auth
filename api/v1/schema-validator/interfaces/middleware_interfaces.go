package interfaces

import "github.com/gin-gonic/gin"

type SchemaValidatorMiddlewareInterfaces interface {
	ValidationRequestMiddleware(schemaMap map[string]string) gin.HandlerFunc
}
