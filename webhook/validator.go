package webhook

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func ValidateStruct(webhookRequest Request) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	validate.RegisterValidation("clientpath", ValidateClientPath)

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

func ValidateClientPath(fl validator.FieldLevel) bool {

	cp := viper.GetString("clientpath")

	if len(fl.Field().String()) < len(cp) {
		return false
	}

	if fl.Field().String()[0:len(cp)] == cp {
		return true
	}
	return false
}
