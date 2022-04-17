package models

import (
	"github.com/cryptocurrency-trading-bot/algos"
)

type DataFrameCandle struct {
	Candles []Candle `json:"candles"`
	Smas    []Sma    `json:"smas,omitempty"`
	Emas    []Ema    `json:"emas,omitempty"`
}

type Sma struct {
	Period int       `json:"period"`
	Values []float64 `json:"values"`
}

type Ema struct {
	Period int       `json:"period"`
	Values []float64 `json:"values"`
}

func (df *DataFrameCandle) Closes() []float64 {
	closes := make([]float64, len(df.Candles))

	for i, candle := range df.Candles {
		closes[i] = candle.Close
	}

	return closes
}

func (df *DataFrameCandle) AddSma(period int) {
	if len(df.Candles) >= period {
		df.Smas = append(df.Smas, Sma{period, algos.CalcSma(period, df.Closes())})
	}
}

func (df *DataFrameCandle) AddEma(period int) {
	if len(df.Candles) >= period {
		df.Emas = append(df.Emas, Ema{period, algos.CalcEma(period, df.Closes())})
	}
}
