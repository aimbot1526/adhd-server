package models

import "time"

type ProductCategory struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Created_At time.Time
	Updated_At time.Time
}

func MapPC(p *ProductCategory) *ProductCategory {
	temp := ProductCategory{
		ID:         p.ID,
		Name:       p.Name,
		Created_At: p.Created_At,
		Updated_At: p.Updated_At,
	}
	return &temp
}

func (p *ProductCategory) Create() error {

	err := GetDB().Create(p).Error

	if err != nil {
		return db.Error
	}
	return nil
}

func (p *ProductCategory) Delete() error {

	err := GetDB().Delete(p).Error

	if err != db.Error {
		return db.Error
	}
	return nil
}

func UpdatePC(p *ProductCategory) error {

	err := GetDB().First(p).Error

	if err != nil {
		return err
	}

	GetDB().Save(p)

	return nil
}

func FindByProductCategoryName(name string) *ProductCategory {

	temp := ProductCategory{Name: name}

	err := GetDB().Where("name = ?").First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}

func FindPcById(id int) *ProductCategory {

	temp := ProductCategory{ID: id}

	err := GetDB().Where("id = ?").First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}
