package controllers

import (
	"time"

	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/gofiber/fiber/v2"
)

func CreateDiscount(c *fiber.Ctx) error {

	d := &models.Discount{}

	if err := c.BodyParser(d); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	exDiscount := models.FindByDiscountName(d.Name)

	if exDiscount.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Discount already exists. Please try again!",
		})
	}

	d.Active = true
	d.Created_At = time.Now()
	d.Updated_At = time.Now()

	e := d.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	resp := models.MapDiscount(d)

	return c.JSON(resp)
}
