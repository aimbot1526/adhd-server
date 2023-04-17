package request

type CreateProductRequest struct {
	Name               string
	Description        string
	Price              int
	ProductCategoryId  int
	ProductInventoryId int
	DiscountId         int
}
