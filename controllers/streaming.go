package controllers

import (
	"github.com/cryptocurrency-trading-bot/models"
)

func Streaming() {
	tickerChannel := make(chan models.Ticker)
	go models.GetRealTimeTicker(tickerChannel)

	go func() {
		for ticker := range tickerChannel {
			models.CreateCandle(&ticker)
		}
	}()
}
