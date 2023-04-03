package models

import "time"

type CartItem struct {
	ID                int `gorm:"primaryKey"`
	Quantity          int
	Created_At        time.Time
	Updated_At        time.Time
	ProductID         int // one-to-one example
	Product           Product
	ShoppingSessionID int // foreign key
	ShoppingSession   ShoppingSession
}
