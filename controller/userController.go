package controller

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/gofiber/fiber/v2"
)

type CreateUserParams struct {
	NamaUser string `json:"namaUser"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if role is not admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Parse body
	u := new(CreateUserParams)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Return if any of the params is empty
	if u.NamaUser == "" || u.Password == "" || u.Role == "" || u.Username == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Return if role is neither admin, manajer, or kasir
	if u.Role != "admin" && u.Role != "manajer" && u.Role != "kasir" {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Bad request: invalid role"))
	}

	// Return if username is already exist
	var users []model.User
	db.DB.Where("username = ?", u.Username).Find(&users)
	if len(users) > 0 {
		return c.Status(fiber.StatusConflict).Send([]byte("Username already exist"))
	}

	// Hash the password with SHA-256
	h := sha256.New()
	h.Write([]byte(u.Password))
	hashedPass := h.Sum(nil)
	hexPass := hex.EncodeToString(hashedPass)

	// Query to database
	newUser := model.User{NamaUser: u.NamaUser, Role: u.Role, Username: u.Username, Password: hexPass}

	// Return the error if any
	err := db.DB.Create(&newUser).Error
	if err != nil {
		return err
	}
	return c.Status(201).Send([]byte("User created"))
}
