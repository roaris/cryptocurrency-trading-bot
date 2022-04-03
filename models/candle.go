package models

import (
	"time"
)

type Candle struct {
	ID        uint
	Open      float64   `gorm:"not null"`
	Close     float64   `gorm:"not null"`
	High      float64   `gorm:"not null"`
	Low       float64   `gorm:"not null"`
	Volume    float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"not null;unique"`
}

func findCandleWithTime(t time.Time) (candle *Candle) {
	err := db.Where("timestamp = ?", t).First(&candle).Error
	if err != nil {
		return nil
	}
	return candle
}

func CreateCandle(t *Ticker) (err error) {
	truncateDateTime := t.truncateDateTime()
	candle := findCandleWithTime(truncateDateTime)
	price := t.getMidPrice()

	if candle == nil {
		err = db.Save(&Candle{Open: price, Close: price, High: price, Low: price, Volume: t.Volume, Timestamp: truncateDateTime}).Error
	} else {
		candle.Close = price
		if candle.High < price {
			candle.High = price
		}
		if candle.Low > price {
			candle.Low = price
		}
		candle.Volume += t.Volume
		err = db.Save(candle).Error
	}
	return err
}
