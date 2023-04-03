package models

import "time"

type OrderDetails struct {
	ID               int `gorm:"primaryKey"`
	Total            int
	Created_At       time.Time
	Updated_At       time.Time
	UserID           int
	User             User
	PaymentDetailsID int
	PaymentDetails   PaymentDetails
}
