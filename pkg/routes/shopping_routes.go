package routes

import (
	"github.com/aimbot1526/adhd-server/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func ShoppingRoute(a *fiber.App) {

	route := a.Group("/api/v1/shop")

	route.Post("/add/session/:id", controllers.CreateShoppingSession)

	route.Patch("/update/session/:id", controllers.UpdateShoppingSession)

	route.Post("/cart/add", controllers.AddProductToCart)

	route.Get("/cart/list", controllers.GetAllCartItems)
}
