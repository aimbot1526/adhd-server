package controllers

import (
	"strconv"
	"time"

	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/gofiber/fiber/v2"
)

func CreateShoppingSession(c *fiber.Ctx) error {

	s := &models.ShoppingSession{}

	if err := c.BodyParser(s); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later! No userId found",
		})
	}

	newSession := models.ShoppingSession{
		Total:      s.Total,
		Created_At: time.Now(),
		Updated_At: time.Now(),
		UserID:     uint(id),
	}

	e := newSession.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	resp := models.MapSession(&newSession)

	return c.JSON(resp)
}

// Update on user exiting the platform
func UpdateShoppingSession(c *fiber.Ctx) error {

	uId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later! No userId found",
		})
	}

	latestSession, e := models.FindLatestShoppingSession(uId)

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later! No userId found",
		})
	}

	updatedSession, er := models.UpdateSession(latestSession, "updated_at", time.Now())

	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Unable to update! Please try again later",
		})
	}

	resp := models.MapSession(updatedSession)

	return c.JSON(resp)
}
