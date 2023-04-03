package models

import "time"

type OrderItems struct {
	ID             int `gorm:"primaryKey"`
	Quantity       int
	Created_At     time.Time
	Updated_At     time.Time
	OrderDetailsID int
	OrderDetails   OrderDetails
	ProductID      int
	Product        Product
}
