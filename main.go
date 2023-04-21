package main

import (
	"github.com/aimbot1526/adhd-server/cmd/server"
	"github.com/aimbot1526/adhd-server/configs"
	"github.com/aimbot1526/adhd-server/pkg/middleware"
	"github.com/aimbot1526/adhd-server/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.ProductRoute(app)

	routes.UserRoute(app)

	routes.ShoppingRoute(app)

	server.StartServer(app)
}
