package validator

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateSchema(schemaJson, bodyJson string) (bool, string, error) {
	schemaLoader := gojsonschema.NewStringLoader(schemaJson)
	documentLoader := gojsonschema.NewStringLoader(bodyJson)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false, "", err
	}

	if result.Valid() {
		return true, "", nil
	}

	var errMsg string
	for _, desc := range result.Errors() {
		errMsg += fmt.Sprintf("- %s\n", desc)
	}
	return false, errMsg, nil
}
