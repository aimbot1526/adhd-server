package routes

import (
	"github.com/aimbot1526/adhd-server/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoute(a *fiber.App) {

	route := a.Group("/api/v1/products")

	route.Post("/add", controllers.CreateProduct)

	route.Post("/add/offer", controllers.CreateDiscount)

	route.Post("/add/category", controllers.CreateProductCategory)

	route.Post("/add/inventory", controllers.CreateProductInventory)
}
