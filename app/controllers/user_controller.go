package controllers

import (
	"time"

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

	exUser := models.FindByEmail(u.Email)

	if exUser.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "User already exists. Please try again!",
		})
	}

	u.Created_At = time.Now()
	u.Updated_At = time.Now()
	u.Role = "ROLE_USER"

	e := u.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	resp := models.MapUser(u)

	return c.JSON(resp)
}
