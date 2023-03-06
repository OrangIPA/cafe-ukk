package controller

import (
	"errors"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateMeja(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return unaouthorized if role is not admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get nomorMeja from request body, check if it is empty, then query it to the database and return error if any
	nomorMeja := c.FormValue("nomorMeja")
	if nomorMeja == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := db.DB.Create(&model.Meja{NomorMeja: nomorMeja}).Error; err != nil {
		return err
	}

	// Return created
	return c.SendStatus(fiber.StatusCreated)
}

func GetAllMeja(c *fiber.Ctx) error {
	// Query to database and return the error if any
	var mejas []model.Meja
	if err := db.DB.Find(&mejas).Error; err != nil {
		return err
	}

	// Return all mejas
	return c.JSON(mejas)
}

func GetMejaById(c *fiber.Ctx) error {
	// Get meja id
	mejaId := c.Params("id")

	// Query to database and return the error if any
	var meja model.Meja
	if err := db.DB.First(&meja, mejaId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return err
	}

	// Return the meja
	return c.JSON(meja)
}