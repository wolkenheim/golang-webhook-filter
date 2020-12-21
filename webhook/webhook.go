package webhook

import (
	"github.com/gofiber/fiber/v2"
)

// Request : describes JSON post body
type Request struct {
	AssetID  string `json:"assetId" validate:"required,min=3,max=32"`
	Metadata struct {
		FolderPath             string `json:"folderPath" validate:"required,clientpath"`
		CfApprovalStateClient1 string `json:"cf_approvalState_client1" validate:"required,eq=Approved|eq=Rejected"`
		CfAssetType            struct {
			Value string `json:"value" validate:"required,eq=Content Image|eq=Product Image"`
		} `json:"cf_assetType" validate:"required,dive"`
	} `json:"metadata" validate:"required,dive"`
}

// ErrorResponse : defines validation error
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// CreateWebhook : route handler for post data
func CreateWebhook(c *fiber.Ctx) error {

	webhookRequest := new(Request)

	bodyParserError := c.BodyParser(webhookRequest)
	if bodyParserError != nil {
		return c.Status(400).JSON(bodyParserError)
	}

	validationErrors := validateStruct(*webhookRequest)
	if validationErrors != nil {
		return c.Status(400).JSON(validationErrors)
	}

	asset := AssetWithStatus{
		AssetID: webhookRequest.AssetID,
		Status:  webhookRequest.Metadata.CfApprovalStateClient1,
	}

	asset.send()

	return c.JSON(webhookRequest)
}
