package entity

import "time"

type Image struct {
	ID        string `gorm:"primary_key"`
	URL       string
	Keyword   string
	UsedCount int  `gorm:"default:0"`
	Reported  bool `gorm:"default:false"`
	Confirmed bool `gorm:"default:false"`
	CreatedAt time.Time
}
