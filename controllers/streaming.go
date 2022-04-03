package controllers

import (
	"encoding/json"
	"time"

	"github.com/cryptocurrency-trading-bot/models"
	"github.com/cryptocurrency-trading-bot/utils"
)

func Streaming() {
	for {
		resp, _ := utils.DoHttpRequest("GET", "https://api.bitflyer.com/v1/ticker", nil, nil, nil)
		var ticker models.Ticker
		json.Unmarshal(resp, &ticker)
		models.CreateCandle(&ticker)
		time.Sleep(time.Second)
	}
}
