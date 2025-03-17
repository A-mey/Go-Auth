package interfaces

type SchemaValidatorServiceInterface interface {
	ValidateSchema(schema string, data interface{}) error
}
