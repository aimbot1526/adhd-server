package models

import "time"

type Discount struct {
	ID              int `gorm:"primaryKey"`
	Name            string
	DiscountPercent float32
	Active          bool
	Created_At      time.Time
	Updated_At      time.Time
}
