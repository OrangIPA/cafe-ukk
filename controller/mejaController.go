package controller

import (
	"errors"
	"strconv"

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

func UpdateMeja(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if insufficient role
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get request id
	mejaIdBef, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	mejaId := uint(mejaIdBef)

	// Return 404 if meja doesn't exist
	var testMeja model.Meja
	if err := db.DB.First(&testMeja, mejaId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return err
	}

	// Perform query
	if err := db.DB.Model(&model.Meja{MejaID: mejaId}).Updates(model.Meja{NomorMeja: c.FormValue("nomorMeja")}).Error; err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

func DeleteMeja(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if insufficient role
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get mejaId, parse it and then return the error if any
	mejaId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Check if meja is exist. if not, return 404
	var meja model.Meja
	if err := db.DB.First(&meja, mejaId).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Query to database and return the error if any
	if err := db.DB.Delete(&model.Meja{}, mejaId).Error; err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}