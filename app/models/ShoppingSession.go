package models

import (
	"time"

	"github.com/aimbot1526/adhd-server/pkg/payload/response"
)

type ShoppingSession struct {
	ID         int `gorm:"primaryKey"`
	Total      int
	Created_At time.Time
	Updated_At time.Time
	UserID     uint // foreign key
}

func MapSession(s *ShoppingSession) *response.ShoppingSessionResponse {

	temp := response.ShoppingSessionResponse{
		Total:      s.Total,
		Created_At: s.Created_At,
		Updated_At: s.Updated_At,
		UserID:     s.UserID,
	}

	return &temp
}

func (s *ShoppingSession) Create() error {

	err := GetDB().Create(s).Error

	if err != nil {
		return db.Error
	}
	return nil
}

func FindLatestShoppingSession(id int) (*ShoppingSession, error) {

	temp := ShoppingSession{}

	err := GetDB().Raw(
		"SELECT DISTINCT * FROM shopping_sessions WHERE user_id = ? ORDER BY updated_at DESC LIMIT 1", id).
		Scan(&temp).Error

	if err != nil {
		return nil, err
	}
	return &temp, nil
}

func UpdateSession(s *ShoppingSession, column string, value interface{}) (*ShoppingSession, error) {

	err := GetDB().Model(&s).Update(column, value).Error

	if err != nil {
		return nil, err
	}
	return s, nil
}
