package controller

import (
	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/gofiber/fiber/v2"
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