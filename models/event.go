package models

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type SignalEvent struct {
	ID        uint      `json:"id"`
	Side      string    `gorm:"not null" json:"side"`
	Size      float64   `gorm:"not null" json:"size"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func (SignalEvent) TableName() string {
	productCode := os.Getenv("PRODUCT_CODE")
	return fmt.Sprintf("%s_signal_events", strings.ToLower(productCode))
}

func NewSignalEvent(side string, size float64) *SignalEvent {
	return &SignalEvent{Side: side, Size: size}
}

func (s *SignalEvent) Save() error {
	err := db.Save(s).Error
	return err
}
