package models

import "time"

type PaymentDetails struct {
	ID         int `gorm:"primaryKey"`
	Amount     int
	Provider   string
	Status     string
	Created_At time.Time
	Updated_At time.Time
}
