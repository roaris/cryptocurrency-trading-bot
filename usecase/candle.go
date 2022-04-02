package usecase

import (
	"time"

	"github.com/cryptocurrency-trading-bot/model"
	"gorm.io/gorm"
)

type Candle struct {
	db *gorm.DB
}

func NewCandle(db *gorm.DB) Candle {
	return Candle{db}
}

func (c *Candle) FindWithTime(time time.Time) *model.Candle {
	var candles []model.Candle
	c.db.Where("timestamp = ?", time).Find(&candles)
	if len(candles) == 0 {
		return nil
	}
	return &candles[0]
}

func (c *Candle) Save(ticker *model.Ticker) error {
	dateTime, err := time.Parse(time.RFC3339, ticker.Timestamp+"Z")
	if err != nil {
		return err
	}

	truncateDateTime := dateTime.Truncate(time.Minute)
	candle := c.FindWithTime(truncateDateTime)
	price := (ticker.BestBid + ticker.BestAsk) / 2

	if candle == nil {
		err = c.db.Save(&model.Candle{
			Open:      price,
			Close:     price,
			High:      price,
			Low:       price,
			Volume:    ticker.Volume,
			Timestamp: truncateDateTime,
		}).Error
	} else {
		candle.Close = price
		if candle.High < price {
			candle.High = price
		}
		if candle.Low > price {
			candle.Low = price
		}
		candle.Volume += ticker.Volume
		err = c.db.Save(&candle).Error
	}

	if err != nil {
		return err
	}
	return nil
}
