package validators

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

func validateMaxLength(fieldLevel validator.FieldLevel) bool {
	currentLength := fieldLevel.Field().Len()
	allowedLength, err := strconv.Atoi(fieldLevel.Param())
	if err != nil {
		return false
	}
	if currentLength > allowedLength {
		return false
	}
	return true
}
