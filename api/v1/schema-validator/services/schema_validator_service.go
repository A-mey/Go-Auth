package services

import "github.com/A-mey/GO-AUTH/api/v1/schema-validator/interfaces"

var _ interfaces.SchemaValidatorServiceInterface = (*SchemaValidatorServices)(nil)

type SchemaValidatorServices struct{}
