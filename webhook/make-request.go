package webhook

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
)

func postRequest(asset Asset){

	jsonValue, _ := json.Marshal(asset)

	resp, err := http.Post("http://localhost:3000/mock-api", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v", resp)
}