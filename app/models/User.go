package models

import "time"

type User struct {
	ID              int `gorm:"primaryKey"`
	Email           string
	Phone           int
	Password        string
	Created_At      time.Time
	Updated_At      time.Time
	Role            string
	UserAddress     []UserAddress
	ShoppingSession []ShoppingSession
	UserPayment     []UserPayment
}
