package routes

import (
	"github.com/aimbot1526/adhd-server/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(a *fiber.App) {

	route := a.Group("/api/v1/user")

	route.Post("/create", controllers.CreateUser)
}
