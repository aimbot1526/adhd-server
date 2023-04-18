package controllers

import (
	"strconv"
	"time"

	"github.com/aimbot1526/adhd-server/app/models"
	"github.com/aimbot1526/adhd-server/pkg/payload/request"
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

	exUser := models.FindByUserEmail(u.Email)

	if exUser.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "User already exists. Please try again!",
		})
	}

	newUser := models.User{
		Email:      u.Email,
		Phone:      u.Phone,
		Password:   u.Password,
		Created_At: time.Now(),
		Updated_At: time.Now(),
		Role:       "ROLE_USER",
	}

	e := newUser.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later !",
		})
	}

	resp := models.MapUser(&newUser)

	return c.JSON(resp)
}

func AddUserAddress(c *fiber.Ctx) error {

	uar := &request.AddUserAddrRequest{}

	if err := c.BodyParser(uar); err != nil {
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

	exAddr := models.FindAddrByUserId(id)

	if exAddr.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Product already exists. Please try again!",
		})
	}

	newAddr := models.UserAddress{
		Address_line1: uar.Address_line1,
		Address_line2: uar.Address_line2,
		City:          uar.City,
		PostalCode:    uar.PostalCode,
		Country:       uar.Country,
		Mobile:        uar.Mobile,
		Created_At:    time.Now(),
		Updated_At:    time.Now(),
		UserID:        uint(id),
	}

	e := newAddr.Create()

	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again later!",
		})
	}

	resp := models.MapUserAddr(&newAddr)

	return c.JSON(resp)
}
