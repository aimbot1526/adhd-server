package request

type AddToCartRequest struct {
	Quantity          int
	ProductId         int
	ShoppingSessionId int
}
