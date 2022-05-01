package models

import "time"

type SignalEvent struct {
	ID        uint      `json:"id"`
	Side      string    `gorm:"not null" json:"side"`
	Size      float64   `gorm:"not null" json:"size"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
