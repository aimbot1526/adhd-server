package models

import "time"

type ProductInventory struct {
	ID         int `gorm:"primaryKey"`
	Quantity   int
	Created_At time.Time
	Updated_At time.Time
}

func MapPI(p *ProductInventory) *ProductInventory {
	temp := ProductInventory{
		ID:         p.ID,
		Quantity:   p.Quantity,
		Created_At: p.Created_At,
		Updated_At: p.Updated_At,
	}
	return &temp
}

func (p *ProductInventory) Create() error {

	err := GetDB().Create(p).Error

	if err != nil {
		return db.Error
	}
	return nil
}

func (p *ProductInventory) Delete() error {

	err := GetDB().Delete(p).Error

	if err != db.Error {
		return db.Error
	}
	return nil
}

func UpdatePI(p *ProductInventory) error {

	err := GetDB().First(p).Error

	if err != nil {
		return err
	}

	GetDB().Save(p)

	return nil
}

func FindByQuantity(quantity int) *ProductInventory {

	temp := ProductInventory{Quantity: quantity}

	err := GetDB().Where("quantity = ?").First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}

func FindPiById(id int) *ProductInventory {

	temp := ProductInventory{ID: id}

	err := GetDB().Where("id = ?").First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}
