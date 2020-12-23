package webhook

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// ValidationErrorResponse : defines validation error
type ValidationErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func validateStruct(webhookRequest Request) []*ValidationErrorResponse {
	var errors []*ValidationErrorResponse
	validate := validator.New()
	validate.RegisterValidation("clientpath", validateClientPath)

	err := validate.Struct(webhookRequest)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func validateClientPath(fl validator.FieldLevel) bool {

	cp := viper.GetString("clientpath")

	if len(fl.Field().String()) < len(cp) {
		return false
	}

	if fl.Field().String()[0:len(cp)] == cp {
		return true
	}
	return false
}
