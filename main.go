package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/aimbot1526/adhd-server/cmd/server"
	"github.com/aimbot1526/adhd-server/configs"
	"github.com/aimbot1526/adhd-server/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func loadEnv() error {
	const projectDirName = "adhd-server"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return err
}

func main() {

	e := loadEnv()
	if e != nil {
		fmt.Print(e)
	}

	config := configs.FiberConfig()

	app := fiber.New(config)

	routes.ProductRoute(app)

	server.StartServer(app)
}
