package models

import "time"

type Product struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Price       int
	Created_At  time.Time
	Updated_At  time.Time
}
