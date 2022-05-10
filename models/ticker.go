package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickId          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (t *Ticker) getMidPrice() float64 {
	return (t.BestAsk + t.BestBid) / 2
}

func (t *Ticker) truncateDateTime() time.Time {
	dateTime, _ := time.Parse(time.RFC3339, t.Timestamp)
	return dateTime.Truncate(time.Minute)
}

type jsonrpc2 struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Result  interface{} `json:"result,omitempty"`
	ID      *int        `json:"id,omitemtpy"`
}

type subscribeParams struct {
	Channel string `json:"channel"`
}

func GetRealTimeTicker(ch chan<- Ticker) {
	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	productCode := os.Getenv("PRODUCT_CODE")
	channel := fmt.Sprintf("lightning_ticker_%s", productCode)
	if err := c.WriteJSON(&jsonrpc2{Version: "2.0", Method: "subscribe", Params: subscribeParams{channel}}); err != nil {
		log.Fatal(err)
	}

OUTER:
	for {
		var message jsonrpc2
		if err := c.ReadJSON(&message); err != nil {
			log.Println(err)
			continue
		}

		if message.Method == "channelMessage" {
			for k, v := range (message.Params).(map[string]interface{}) {
				if k == "message" {
					marshalTicker, err := json.Marshal(v)
					if err != nil {
						continue OUTER
					}
					var ticker Ticker
					err = json.Unmarshal(marshalTicker, &ticker)
					if err != nil {
						continue OUTER
					}
					ch <- ticker
				}
			}
		}
	}
}
