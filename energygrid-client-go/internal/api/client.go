package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"fmt"


	"energygrid-client-go/internal/utils"
)

const (
	APIURL  = "http://localhost:3000/device/real/query"
	APIPath = "/device/real/query"
	Token   = "interview_token_123"
)

var (
	ErrRateLimited = errors.New("rate limited")
)

type RequestBody struct {
	SNList []string `json:"sn_list"`
}

type DeviceData struct {
	SN          string `json:"sn"`
	Power       string `json:"power"`
	Status      string `json:"status"`
	LastUpdated string `json:"last_updated"`
}

type APIResponse struct {
	Data []DeviceData `json:"data"`
}

// Fetch a single batch (max 10 SNs)
func FetchBatch(snList []string) ([]DeviceData, error) {
	body := RequestBody{SNList: snList}
	jsonBody, _ := json.Marshal(body)

	timestamp := time.Now().UnixMilli()
	signature := utils.GenerateSignature(
		APIPath,
		Token,
		fmt.Sprintf("%d", timestamp),
	)

	req, err := http.NewRequest("POST", APIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("timestamp", fmt.Sprintf("%d", timestamp))
	req.Header.Set("signature", signature)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return nil, ErrRateLimited
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var apiResp APIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return nil, err
	}

	return apiResp.Data, nil
}
