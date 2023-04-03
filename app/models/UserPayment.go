package models

import "time"

type UserPayment struct {
	ID          int `gorm:"primaryKey"`
	PaymentType string
	Provider    string
	AccountNo   int
	Expiry      time.Time
	Created_At  time.Time
	Updated_At  time.Time
	UserID      int // foreign key
}
