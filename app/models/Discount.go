package models

import "time"

type Discount struct {
	ID              int `gorm:"primaryKey"`
	Name            string
	DiscountPercent float32
	Active          bool
	Created_At      time.Time
	Updated_At      time.Time
}

func MapDiscount(d *Discount) *Discount {
	temp := Discount{
		ID:              d.ID,
		Name:            d.Name,
		Active:          d.Active,
		Created_At:      d.Created_At,
		Updated_At:      d.Updated_At,
		DiscountPercent: d.DiscountPercent,
	}
	return &temp
}

func (discount *Discount) Create() error {

	err := GetDB().Create(discount).Error

	if err != nil {
		return db.Error
	}
	return nil
}

func (d *Discount) Delete() error {

	err := GetDB().Delete(d).Error

	if err != db.Error {
		return db.Error
	}
	return nil
}

func UpdateDiscount(d *Discount) error {

	err := GetDB().First(d).Error

	if err != nil {
		return err
	}

	GetDB().Save(d)

	return nil
}

func FindByDiscountName(name string) *Discount {

	temp := Discount{Name: name}

	err := GetDB().Where("name = ?").First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}

func FindDiscountById(id int) *Discount {

	temp := Discount{ID: id}

	err := GetDB().Where("id = ?").First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}
