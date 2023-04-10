package controllers

import (
	"time"

	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProductCategory(c *fiber.Ctx) error {

	pc := &models.ProductCategory{}

	if err := c.BodyParser(pc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	exCategory := models.FindByProductCategoryName(pc.Name)

	if exCategory.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Category already exists. Please try again!",
		})
	}

	pc.Created_At = time.Now()
	pc.Updated_At = time.Now()

	e := pc.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	resp := models.MapPC(pc)

	return c.JSON(resp)
}
