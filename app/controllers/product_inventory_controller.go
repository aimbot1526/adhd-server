package controllers

import (
	"time"

	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProductInventory(c *fiber.Ctx) error {

	pi := &models.ProductInventory{}

	if err := c.BodyParser(pi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	exCategory := models.FindByQuantity(pi.Quantity)

	if exCategory.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Discount already exists. Please try again!",
		})
	}

	pi.Created_At = time.Now()
	pi.Updated_At = time.Now()

	e := pi.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	resp := models.MapPI(pi)

	return c.JSON(resp)
}
