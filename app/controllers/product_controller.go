package controllers

import (
	"time"

	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {

	p := &models.Product{}

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	exProduct := models.FindByProductName(p.Name)

	if exProduct.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Discount already exists. Please try again!",
		})
	}

	pc := models.FindPcById(p.ProductCategoryID)

	if pc.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "No category found. Please try again!",
		})
	}

	pi := models.FindPiById(p.ProductInventoryID)

	if pc.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "No stock found. Please try again!",
		})
	}

	d := models.FindDiscountById(p.DiscountID)

	if d.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "No category found. Please try again!",
		})
	}

	p.Created_At = time.Now()
	p.Updated_At = time.Now()
	p.ProductCategory = *pc
	p.ProductInventory = *pi
	p.Discount = *d

	e := p.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	resp := models.MapProduct(p)

	return c.JSON(resp)
}
