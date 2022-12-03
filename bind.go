package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Function To Bind Validators
func BindValidators() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// registering validator for Max Length
		validate.RegisterValidation("maxlength", validateMaxLength)
	}
}
