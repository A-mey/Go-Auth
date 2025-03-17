package interfaces

type SchemaValidatorServiceInterface interface {
	ValidateSchema(key, body string, schemaMap map[string]string) (bool, string, error)
}
