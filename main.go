package main

import (
	"encoding/json"
	"time"

	"github.com/cryptocurrency-trading-bot/usecase"

	"github.com/cryptocurrency-trading-bot/model"
	"github.com/cryptocurrency-trading-bot/repository"
	"github.com/cryptocurrency-trading-bot/utils"
)

func main() {
	candleUseCase := usecase.NewCandle(repository.DB)

	for {
		resp, _ := utils.DoHttpRequest("GET", "https://api.bitflyer.com/v1/ticker", nil, nil, nil)
		var ticker model.Ticker
		json.Unmarshal(resp, &ticker)
		candleUseCase.Save(&ticker)
		time.Sleep(time.Second)
	}
}
