package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type CreateUserParams struct {
	NamaUser string `form:"namaUser"`
	Role     string `form:"role"`
	Username string `form:"username"`
	Password string `form:"password"`
}

type UpdateUserParams struct {
	UserID   uint   `form:"userId"`
	NamaUser string `form:"namaUser"`
	Role     string `form:"role"`
	Username string `form:"username"`
	Password string `form:"password"`
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
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return if any of the params is empty
	if u.NamaUser == "" || u.Password == "" || u.Role == "" || u.Username == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Return if role is neither admin, manajer, or kasir
	if u.Role != "admin" && u.Role != "manager" && u.Role != "kasir" {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Bad request: invalid role"))
	}

	// Return if username is already exist
	var users []model.User
	if err := db.DB.Where("username = ?", u.Username).Find(&users).Error; err != nil {
		return err
	}
	if len(users) > 0 {
		return c.Status(fiber.StatusConflict).Send([]byte("Username already exist"))
	}

	// Hash the password with SHA-256
	h := sha256.New()
	h.Write([]byte(u.Password))
	hashedPass := h.Sum(nil)
	hexPass := hex.EncodeToString(hashedPass)

	// Create entry model
	newUser := model.User{NamaUser: u.NamaUser, Role: u.Role, Username: u.Username, Password: hexPass}

	// Create the entry and return the error if any
	err := db.DB.Create(&newUser).Error
	if err != nil {
		return err
	}
	return c.Status(201).Send([]byte("User created"))
}

func GetAllUser(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if role is not admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Query to database
	var users []model.User
	if err := db.DB.Find(&users).Error; err != nil {
		return err
	}

	// Return the users
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return of role isn't admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Query to database
	var user model.User
	if err := db.DB.First(&user, c.Params("id")).Error; err != nil {
		return err
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if role is not admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Parse body
	user := new(UpdateUserParams)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	// Return if any of the params is empty
	if user.UserID == 0 || user.NamaUser == "" || user.Password == "" || user.Role == "" || user.Username == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Return if role is neither admin, manajer, or kasir
	if user.Role != "admin" && user.Role != "manager" && user.Role != "kasir" {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Bad request: invalid role"))
	}

	// Hash the password with SHA-256
	h := sha256.New()
	h.Write([]byte(user.Password))
	hashedPass := h.Sum(nil)
	hexPass := hex.EncodeToString(hashedPass)

	// Create entry model
	u := model.User{
		UserID:   user.UserID,
		NamaUser: user.NamaUser,
		Role:     user.Role,
		Username: user.Username,
		Password: hexPass,
	}

	// Update the menu and return error if any
	if err := db.DB.Save(&u).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusBadRequest).SendString("nothing is changed")
		}
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteUser(c *fiber.Ctx) error {
	// Get token claims
	claims := helper.TokenClaims(c)
	role := claims["role"].(string)

	// Return if role is not admin
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get UserId, parse it into int, and return 400 if failed to parse
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Check if menu is exist. if not, return 404
	var user model.User
	if err := db.DB.First(&user, userId).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Query to database and return the error if any
	if err := db.DB.Delete(&model.User{}, userId).Error; err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
