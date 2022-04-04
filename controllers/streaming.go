package controllers

import (
	"github.com/cryptocurrency-trading-bot/models"
)

func Streaming() {
	tickerChannel := make(chan models.Ticker)
	go models.GetRealTimeTicker(tickerChannel)

	for ticker := range tickerChannel {
		models.CreateCandle(&ticker)
	}
}
