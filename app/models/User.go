package models

import (
	"time"
)

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

func (user *User) Create() error {

	err := GetDB().Create(user).Error

	if err != nil {
		return db.Error
	}
	return nil
}

func View(id int) *User {

	temp := User{ID: id}

	err := GetDB().First(&temp, "id = ?", id).Error

	if err != nil {
		return &temp
	}
	return &temp
}
