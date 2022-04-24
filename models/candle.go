package models

import (
	"time"
)

type Candle struct {
	ID        uint      `json:"id"`
	Open      float64   `gorm:"not null" json:"open"`
	Close     float64   `gorm:"not null" json:"close"`
	High      float64   `gorm:"not null" json:"high"`
	Low       float64   `gorm:"not null" json:"low"`
	Volume    float64   `gorm:"not null" json:"volume"`
	Timestamp time.Time `gorm:"not null;unique" json:"timestamp"`
}

func findCandleWithTime(t time.Time) (candle *Candle) {
	err := db.Where("timestamp = ?", t).First(&candle).Error
	if err != nil {
		return nil
	}
	return candle
}

func CreateCandle(t *Ticker) (err error) {
	truncateDateTimeUTC := t.truncateDateTime()
	truncateDateTimeJST := truncateDateTimeUTC.Add(9 * time.Hour)
	candle := findCandleWithTime(truncateDateTimeJST)
	price := t.getMidPrice()

	if candle == nil {
		err = db.Save(&Candle{Open: price, Close: price, High: price, Low: price, Volume: t.Volume, Timestamp: truncateDateTimeJST}).Error
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

func GetCandles(limit int) (candles []Candle, err error) {
	if db.Table("(?) as c", db.Order("timestamp desc").Limit(limit).Model(&Candle{})).Order("timestamp asc").Find(&candles).Error != nil {
		return nil, err
	}
	return candles, nil
}
