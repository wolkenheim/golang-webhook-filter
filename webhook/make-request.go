package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// AssetWithStatus : defines one exportable entity
type AssetWithStatus struct {
	AssetID string `json:"assetId"`
	Status  string `json:"status"`
}

// AssetWithStatusClient : all clients should have this method
type AssetWithStatusClient interface {
	send()
}

// Send : fire and forget request to external API endpoint
func (a AssetWithStatus) send() {

	jsonValue, _ := json.Marshal(a)

	_, err := http.Post(viper.GetString("api-endpoint"), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("%v", err)
	}
}
