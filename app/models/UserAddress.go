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

func MapUserAddr(ua *UserAddress) *UserAddress {
	temp := UserAddress{
		Address_line1: ua.Address_line1,
		Address_line2: ua.Address_line2,
		City:          ua.City,
		PostalCode:    ua.PostalCode,
		Country:       ua.Country,
		Mobile:        ua.Mobile,
		Created_At:    ua.Created_At,
		Updated_At:    ua.Updated_At,
		UserID:        ua.UserID,
	}
	return &temp
}

func (ua *UserAddress) Create() error {

	err := GetDB().Create(ua).Error

	if err != nil {
		return db.Error
	}
	return nil
}

func FindAddrByUserId(id int) *UserAddress {

	temp := UserAddress{}

	err := GetDB().Where("user_id = ?", id).First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}
