package models

import "time"

type ProductCategory struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Created_At time.Time
	Updated_At time.Time
}
