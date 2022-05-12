package controllers

import (
	"github.com/cryptocurrency-trading-bot/algos"
	"github.com/cryptocurrency-trading-bot/models"
	"github.com/cryptocurrency-trading-bot/utils"
)

type bot struct {
	apiClient *utils.APIClient
	status    string
	busy      bool
}

func (b *bot) trade() {
	if b.busy {
		return
	}
	b.busy = true
	defer func() {
		b.busy = false
	}()

	limit := 15
	candles, err := models.GetCandles(limit)
	if err != nil || len(candles) < limit {
		return
	}

	df := models.DataFrameCandle{Candles: candles}
	emaPeriod1 := 5
	emaPeriod2 := 10
	ema1 := algos.CalcEma(emaPeriod1, df.Closes())
	ema2 := algos.CalcEma(emaPeriod2, df.Closes())

	if ema1[limit-2] < ema2[limit-2] && ema1[limit-1] > ema2[limit-1] && b.status == "SELL" {
		b.status = "BUY"
		b.apiClient.SendOrder("BUY", 0.001)
		signalEvent := models.NewSignalEvent("BUY", 0.001)
		signalEvent.Save()
	}

	if ema1[limit-2] > ema2[limit-2] && ema1[limit-1] < ema2[limit-1] && b.status == "BUY" {
		b.status = "SELL"
		b.apiClient.SendOrder("SELL", 0.001)
		signalEvent := models.NewSignalEvent("SELL", 0.001)
		signalEvent.Save()
	}
}
