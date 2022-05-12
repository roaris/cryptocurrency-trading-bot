package controllers

import (
	"os"

	"github.com/cryptocurrency-trading-bot/models"
	"github.com/cryptocurrency-trading-bot/utils"
)

func Streaming() {
	tickerChannel := make(chan models.Ticker)
	apiClient := utils.NewAPIClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	bot := bot{apiClient, "SELL", false}
	go func() {
		for {
			models.GetRealTimeTicker(tickerChannel)
		}
	}()

	go func() {
		for ticker := range tickerChannel {
			models.CreateCandle(&ticker)
			bot.trade()
		}
	}()
}
