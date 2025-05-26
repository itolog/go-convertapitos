package validation

import (
	"github.com/go-playground/validator/v10"
)

type ErrorFields struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

type XValidator struct {
	Validator *validator.Validate
}

var validate = validator.New()

func NewValidator() *XValidator {
	return &XValidator{
		Validator: validate,
	}
}

func (v XValidator) Validate(data any) []ErrorFields {
	var validationErrors []ErrorFields

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorFields

			elem.Field = err.Field()
			elem.Tag = err.Tag()
			elem.Param = err.Param()

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
