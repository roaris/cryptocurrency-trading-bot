package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

const baseURL = "https://api.bitflyer.com"

type APIClient struct {
	key    string
	secret string
}

type Balance struct {
	CurrentCode string  `json:"currency_code"`
	Amount      float64 `json:"amount"`
	Available   float64 `json:"available"`
}

func NewAPIClient(key, secret string) *APIClient {
	return &APIClient{key, secret}
}

func (a APIClient) getHeader(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	// HMACで電子署名を行う API-SECRETを秘密鍵にする
	msg := timestamp + method + endpoint + string(body)
	mac := hmac.New(sha256.New, []byte(a.secret))
	mac.Write([]byte(msg))
	sign := hex.EncodeToString(mac.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":       a.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

// 資産残高を取得する
func (a APIClient) GetMeBalance() ([]Balance, error) {
	method := "GET"
	endpoint := "/v1/me/getbalance"
	header := a.getHeader(method, endpoint, nil)
	resp, err := DoHttpRequest(method, baseURL+endpoint, header, nil, nil)

	if err != nil {
		return nil, err
	}
	var balance []Balance
	if err := json.Unmarshal(resp, &balance); err != nil {
		return nil, err
	}
	return balance, nil
}
