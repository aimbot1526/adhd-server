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

func FindById(id int) *User {

	temp := User{ID: id}

	err := GetDB().First(&temp, "id = ?", id).Error

	if err != nil {
		return &temp
	}
	return &temp
}

func (user *User) Delete() error {

	err := GetDB().Delete(user).Error

	if err != db.Error {
		return db.Error
	}
	return nil
}

func Update(user *User) error {

	err := GetDB().First(user).Error

	if err != nil {
		return err
	}

	GetDB().Save(user)

	return nil
}

func FindByEmail(email string) *User {

	temp := User{Email: email}

	err := GetDB().Where("email = ?").First(&temp).Error

	if err != nil {
		return &temp
	}

	return &temp
}
