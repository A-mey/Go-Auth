package services

func Validator(key, body string, schemaMap map[string]string) (bool, string, error) {
	schema, exists := schemaMap[key]
	if exists {
		isValid, validationError, err := ValidateSchema(schema, body)
		if err != nil {
			return false, "", err
		}
		return isValid, validationError, nil
	}
	return true, "", nil

}
