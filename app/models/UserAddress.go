package models

import "time"

type UserAddress struct {
	ID            int `gorm:"primaryKey"`
	Address_line1 string
	Address_line2 string
	City          string
	PostalCode    int
	Country       string
	Mobile        int
	Created_At    time.Time
	Updated_At    time.Time
	UserID        uint // foreign key
}
