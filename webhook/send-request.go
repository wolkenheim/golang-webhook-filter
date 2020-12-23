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

// AssetHTTP : send Asset via http protocol
type AssetHTTP struct{}

// Send : fire and forget request to external API endpoint
func (a AssetHTTP) Send(asset *AssetWithStatus) {

	jsonValue, _ := json.Marshal(asset)

	_, err := http.Post(viper.GetString("api-endpoint"), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("%v", err)
	}
}
