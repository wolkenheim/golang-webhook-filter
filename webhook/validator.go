package webhook

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func validateStruct(webhookRequest Request) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	validate.RegisterValidation("clientpath", validateClientPath)

	err := validate.Struct(webhookRequest)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func validateClientPath(fl validator.FieldLevel) bool {

	if len(fl.Field().String()) < 40 {
		return false
	}

	if fl.Field().String()[0:40] == viper.Get("clientpath") {
		return true
	}
	return false
}
