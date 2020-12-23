package webhook

import (
	"dam-webhook/application"
	"encoding/json"
	"net/http"
)

// Request : describes JSON post body
type Request struct {
	AssetID  string `json:"assetId" validate:"required,min=3,max=32"`
	Metadata struct {
		FolderPath             string `json:"folderPath" validate:"required,clientpath"`
		CfApprovalStateClient1 string `json:"cf_ApprovalState_client1" validate:"required,eq=Approved|eq=Rejected"`
		CfAssetType            struct {
			Value string `json:"value" validate:"required,eq=Content Image|eq=Product Image"`
		} `json:"cf_assetType" validate:"required,dive"`
	} `json:"metadata" validate:"required,dive"`
}

// Controller defines Controller to process webhook
type Controller struct {
	App         *application.Application
	AssetClient AssetClient
}

// CreateWebhook : route handler for post data
func (con *Controller) CreateWebhook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		con.App.ClientError(w, application.ErrorResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	xHookHeader := r.Header.Get("X-Hook-Signature")
	if len(xHookHeader) == 0 {
		con.App.ClientError(w, application.ErrorResponse{
			Status:  http.StatusOK,
			Message: "Invalid request",
		})
		return
	}

	contentType := r.Header.Get("Content-Type")
	if len(contentType) == 0 || contentType != "application/json" {
		con.App.ClientError(w, application.ErrorResponse{
			Status:  http.StatusOK,
			Message: "Invalid request",
		})
		return
	}

	if r.ContentLength < 1 {
		con.App.ClientError(w, application.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Body missing",
		})
		return
	}

	var webhookRequest Request
	err := json.NewDecoder(r.Body).Decode(&webhookRequest)
	if err != nil {
		con.App.ClientError(w, application.ErrorResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	validationErrors := validateStruct(webhookRequest)
	if validationErrors != nil {

		raw, _ := json.Marshal(validationErrors)

		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(raw)
		return
	}

	asset := &AssetWithStatus{
		AssetID: webhookRequest.AssetID,
		Status:  webhookRequest.Metadata.CfApprovalStateClient1,
	}

	con.AssetClient.Send(asset)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "request accepted"}`))
	return

}
