package models

import "time"

type ShoppingSession struct {
	ID         int `gorm:"primaryKey"`
	Total      int
	Created_At time.Time
	Updated_At time.Time
	UserID     uint // foreign key
}
