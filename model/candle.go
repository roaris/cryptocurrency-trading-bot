package model

import "time"

type Candle struct {
	ID        uint
	Open      float64   `gorm:"not null"`
	Close     float64   `gorm:"not null"`
	High      float64   `gorm:"not null"`
	Low       float64   `gorm:"not null"`
	Volume    float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}
