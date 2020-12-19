package webhook

import (
	"net/http"
	"encoding/json"
	"github.com/spf13/viper"
	"bytes"
	"fmt"
)

type AssetWithStatus struct {
	AssetId string
	Status string
}

// fire and forget request to external API endpoint
func postRequest(asset AssetWithStatus){

	jsonValue, _ := json.Marshal(asset)

	_, err := http.Post(viper.GetString("api-endpoint"), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("%v", err)
	}
}