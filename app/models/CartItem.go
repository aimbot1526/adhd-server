package models

import (
	"time"

	"github.com/aimbot1526/adhd-server/pkg/payload/response"
)

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

func MapCart(p *response.ProductResponse, c *CartItem) *response.CartItem {
	temp := response.CartItem{
		Quantity:        c.Quantity,
		Created_At:      c.Created_At,
		Updated_At:      c.Updated_At,
		ProductResponse: *p,
	}
	return &temp
}

func (c *CartItem) Create() error {

	err := GetDB().Create(c).Error

	if err != nil {
		return db.Error
	}
	return nil
}

func FindAllItemsInCart() (*[]CartItem, error) {

	temp := []CartItem{}

	err := GetDB().
		Preload("Product").
		Preload("ShoppingSession").
		Find(&temp).Error

	if err != nil {
		return &temp, err
	}
	return &temp, nil
}
