package webhook

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseStruct struct {
	Message string
	Status  uint8
}

type WebhookRequest struct {
	AssetId  string `json:"assetId" validate:"required,min=3,max=32"`
	Metadata struct {
		FolderPath               string `validate:"required,clientpath"`
		Cf_approvalState_client1 string `validate:"required,eq=Approved|eq=Rejected"`
		Cf_assetType             struct {
			Value string `validate:"required,eq=Content Image|eq=Product Image"`
		} `validate:"required,dive"`
	} `json:"metadata" validate:"required,dive"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
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

	asset := AssetWithStatus{
		AssetId: webhookRequest.AssetId,
		Status: webhookRequest.Metadata.Cf_approvalState_client1,
	}

	postRequest(asset)

	return c.JSON(webhookRequest)
}