package controllers

import (
	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	u := &models.User{}

	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	checkExistingUser := models.FindByEmail(u.Email)

	if checkExistingUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "User already exists. Please try again!",
		})
	}

	resp := u.Create()

	if resp != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	return c.JSON(resp)
}
