package models

import "time"

type SignalEvent struct {
	ID        uint      `json:"id"`
	Side      string    `gorm:"not null" json:"side"`
	Size      float64   `gorm:"not null" json:"size"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func NewSignalEvent(side string, size float64) *SignalEvent {
	return &SignalEvent{Side: side, Size: size}
}

func (s *SignalEvent) Save() error {
	err := db.Save(s).Error
	return err
}
