package webhook

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type ResponseStruct struct {
	Message string
	Status  uint8
}

type WebhookRequest struct {
	AssetId  string `validate:"required,min=3,max=32"`
	Metadata struct {
		FolderPath               string `validate:"required,clientpath"`
		Cf_approvalState_client1 string `validate:"required,eq=Approved|eq=Rejected"`
		Cf_assetType             struct {
			Value string `validate:"required,eq=Content Image|eq=Product Image"`
		} `validate:"required,dive"`
	} `validate:"required,dive"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Asset struct {
	AssetId string
}




func CreateWebhook(c *fiber.Ctx) error {

	webhookRequest := new(WebhookRequest)

	bodyParserError := c.BodyParser(webhookRequest)
	if bodyParserError != nil {
		return c.Status(400).JSON(bodyParserError)
	}

	validationErrors := validateStruct(*webhookRequest)
	if validationErrors != nil {
		return c.Status(400).JSON(validationErrors)
	}

	asset := Asset{AssetId:webhookRequest.AssetId}

	postRequest(asset)

	return c.JSON(webhookRequest)
}

func validateStruct(webhookRequest WebhookRequest) []*ErrorResponse {
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