package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
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

type Order struct {
	ProductCode    string  `json:"product_code"`     // BTC_JPY
	ChildOrderType string  `json:"child_order_type"` // LIMIT or MARKET
	Side           string  `json:"side"`             // BUY or SELL
	Size           float64 `json:"size"`
	MinuteToExpire int     `json:"minute_to_expire"` // 期限切れまでの時間
	TimeInForce    string  `json:"time_in_force"` // 執行数量条件 GTC or IOC or FOK
}

type OrderRes struct {
	ChildOrderAcceptanceId string `json:"child_order_acceptance_id"`
}

type OrderStatus struct {
	ChildOrderState string `json:"child_order_state"`
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

// 成行注文を行う
func (a APIClient) SendOrder(side string, size float64) (*OrderRes, error) {
	method := "POST"
	endpoint := "/v1/me/sendchildorder"
	body := Order{"BTC_JPY", "MARKET", side, size, 43200, "GTC"}
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	header := a.getHeader(method, endpoint, data)
	resp, err := DoHttpRequest(method, baseURL+endpoint, header, nil, data)

	if err != nil {
		return nil, err
	}
	var orderRes OrderRes
	if err := json.Unmarshal(resp, &orderRes); err != nil {
		return nil, err
	}
	if orderRes.ChildOrderAcceptanceId == "" {
		return nil, errors.New(string(resp))
	}
	return &orderRes, nil
}

// 注文の状態を取得する
func (a APIClient) GetOrderStatus(childOrderAcceptanceId string) (string, error) {
	method := "GET"
	endpoint := "/v1/me/getchildorders"
	header := a.getHeader(method, endpoint, nil)
	query := map[string]string{
		"child_order_acceptance_id": childOrderAcceptanceId,
	}
	resp, err := DoHttpRequest(method, baseURL+endpoint, header, query, nil)

	if err != nil {
		return "", err
	}
	var orderStatus []OrderStatus
	if err := json.Unmarshal(resp, &orderStatus); err != nil {
		return "", err
	}
	if len(orderStatus) == 0 {
		return "", errors.New("Specified order does not exist.")
	}
	return orderStatus[0].ChildOrderState, nil
}
