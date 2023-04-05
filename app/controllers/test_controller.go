package controllers

import "github.com/gofiber/fiber/v2"

func Testing(c *fiber.Ctx) error {

	data := []struct {
		data string
	}{
		{data: "Hello world"},
	}

	return c.JSON(data)
}
