package models

import (
	"time"

	"github.com/aimbot1526/adhd-server/pkg/payload/response"
)

type Product struct {
	ID                 int `gorm:"primaryKey"`
	Name               string
	Description        string
	Price              int
	Created_At         time.Time
	Updated_At         time.Time
	ProductCategoryID  int
	ProductCategory    ProductCategory
	ProductInventoryID int
	ProductInventory   ProductInventory
	DiscountID         int
	Discount           Discount
}

func MapProduct(p *Product) *response.ProductResponse {
	temp := response.ProductResponse{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Created_At:  p.Created_At,
		Updated_At:  p.Updated_At,
	}
	return &temp
}

func (p *Product) Create() error {

	err := GetDB().Create(p).Error

	if err != nil {
		return db.Error
	}
	return nil
}

func (p *Product) Delete() error {

	err := GetDB().Delete(p).Error

	if err != db.Error {
		return db.Error
	}
	return nil
}

func Update(d *Product) error {

	err := GetDB().First(d).Error

	if err != nil {
		return err
	}

	GetDB().Save(d)

	return nil
}

func FindByProductName(name string) *Product {

	temp := Product{}

	err := GetDB().Where("name = ?", name).First(&temp)

	if err != nil {
		return &temp
	}

	return &temp
}

func GetAllProducts() (*[]Product, error) {

	temp := []Product{}

	err := GetDB().
		Preload("ProductCategory").
		Preload("ProductInventory").
		Preload("Discount").
		Find(&temp).Error

	if err != nil {
		return &temp, err
	}
	return &temp, nil
}

func FindProductById(id int) (*Product, error) {

	temp := Product{}

	err := GetDB().Where("id = ?", id).First(&temp).Error

	if err != nil {
		return &temp, err
	}

	return &temp, nil
}
