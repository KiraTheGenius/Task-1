package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	Origin      string `gorm:"size:255"`
	Destination string `gorm:"size:255"`
	StartTime   time.Time
	EndTime     time.Time
	Aircraft    string `gorm:"size:255"`
}
