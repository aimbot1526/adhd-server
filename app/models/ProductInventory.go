package models

import "time"

type ProductInventory struct {
	ID         int `gorm:"primaryKey"`
	Quantity   int
	Created_At time.Time
	Updated_At time.Time
}
