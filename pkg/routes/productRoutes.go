package routes

import (
	"github.com/aimbot1526/adhd-server/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoute(a *fiber.App) {

	route := a.Group("/api/v1")

	route.Post("/products", controllers.Testing)
}
