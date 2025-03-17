package services

import "github.com/A-mey/GO-AUTH/libs/validator"

func (sv *SchemaValidatorServices) ValidateSchema(key, body string, schemaMap map[string]string) (bool, string, error) {
	schema, exists := schemaMap[key]
	if exists {
		isValid, validationError, err := validator.ValidateSchema(schema, body)
		if err != nil {
			return false, "", err
		}
		return isValid, validationError, nil
	}
	return true, "", nil
}
