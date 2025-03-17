package middlewares

import "github.com/A-mey/GO-AUTH/api/v1/schema-validator/interfaces"

var _ interfaces.SchemaValidatorMiddlewareInterfaces = (*SchemaValidatorMiddleware)(nil)

type SchemaValidatorMiddleware struct {
	schemaValidatorService interfaces.SchemaValidatorServiceInterface
}

func NewSchemaValidatorMiddleware(service interfaces.SchemaValidatorServiceInterface) *SchemaValidatorMiddleware {
	return &SchemaValidatorMiddleware{schemaValidatorService: service}
}
