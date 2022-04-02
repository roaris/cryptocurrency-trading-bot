package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cryptocurrency-trading-bot/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickId          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

type Candle struct {
	ID        uint
	Open      float64   `gorm:"not null"`
	Close     float64   `gorm:"not null"`
	High      float64   `gorm:"not null"`
	Low       float64   `gorm:"not null"`
	Volume    float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}

func main() {
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUserName, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Candle{})

	for {
		resp, _ := utils.DoHttpRequest("GET", "https://api.bitflyer.com/v1/ticker", nil, nil, nil)
		var ticker Ticker
		json.Unmarshal(resp, &ticker)

		price := (ticker.BestBid + ticker.BestAsk) / 2
		dateTime, err := time.Parse(time.RFC3339, ticker.Timestamp+"Z")
		if err != nil {
			log.Print(err)
		}
		truncateDateTime := dateTime.Truncate(time.Minute)

		var candles []Candle
		db.Where("timestamp = ?", truncateDateTime).Find(&candles)

		if len(candles) == 0 {
			err = db.Save(&Candle{
				Open:      price,
				Close:     price,
				High:      price,
				Low:       price,
				Volume:    ticker.Volume,
				Timestamp: truncateDateTime,
			}).Error
		} else {
			candle := candles[0]
			candle.Close = price
			if candle.High < price {
				candle.High = price
			}
			if candle.Low > price {
				candle.Low = price
			}
			candle.Volume += ticker.Volume
			err = db.Save(&candle).Error
		}
		if err != nil {
			log.Print(err)
		}

		time.Sleep(time.Second)
	}
}
