package validator

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
)

func Validator(c *gin.Context, schemaMap map[string]string) (bool, string, error) {
	// key := CreateKey(c)
	body, err := c.GetRawData()
	if err != nil {
		return false, "", err
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	schema, exists := schemaMap[key]
	if exists {
		isValid, validationError, err := ValidateSchema(schema, string(body))
		if err != nil {
			return false, "", err
		}
		return isValid, validationError, nil
	}
	return true, "", nil
}
