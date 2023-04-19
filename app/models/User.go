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

func MapUser(u *User) *User {
	temp := User{
		Email:      u.Email,
		Phone:      u.Phone,
		Created_At: u.Created_At,
		Updated_At: u.Updated_At,
		Role:       u.Role,
	}
	return &temp
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

// TODO : Fix it, example at shoppingSession line 53
// func UpdateUser(user *User) error {

// 	err := GetDB().First(user).Error

// 	if err != nil {
// 		return err
// 	}

// 	GetDB().Save(user)

// 	return nil
// }

func FindByUserEmail(email string) *User {

	temp := User{}

	err := GetDB().Where("email = ?", email).First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}
