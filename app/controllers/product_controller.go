package controllers

import (
	"strconv"
	"time"

	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/aimbot1526/adhd-server/pkg/payload/request"
	"github.com/aimbot1526/adhd-server/pkg/payload/response"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {

	p := &request.CreateProductRequest{}

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
			"msg":   "Product already exists. Please try again!",
		})
	}

	pc := models.FindPcById(p.ProductCategoryId)

	if pc.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "No category found. Please try again!",
		})
	}

	pi := models.FindPiById(p.ProductInventoryId)

	if pc.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "No stock found. Please try again!",
		})
	}

	d := models.FindDiscountById(p.DiscountId)

	if d.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "No category found. Please try again!",
		})
	}

	newProduct := models.Product{
		Name:             p.Name,
		Description:      p.Description,
		Price:            p.Price,
		ProductCategory:  *pc,
		ProductInventory: *pi,
		Discount:         *d,
		Created_At:       time.Now(),
		Updated_At:       time.Now(),
	}

	e := newProduct.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	resp := models.MapProduct(&newProduct)

	return c.JSON(resp)
}

func GetProduct(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later! No userId found",
		})
	}

	p, e := models.FindProductById(id)

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later!",
		})
	}

	resp := models.MapProduct(p)

	return c.JSON(resp)
}

func FindAllProducts(c *fiber.Ctx) error {

	all, err := models.GetAllProducts()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	var filtered []response.ProductResponse

	for _, p := range *all {
		temp := models.MapProduct(&p)
		filtered = append(filtered, *temp)
	}

	return c.JSON(filtered)
}
